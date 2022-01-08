package app

import (
	"encoding/hex"
	"log"

	"github.com/blukat29/evm-explorer/network"
	"github.com/blukat29/evm-explorer/storage"
	"github.com/blukat29/evm-explorer/util"
)

// extAddr -> extCodeID
const addrTable = "addr"

func FetchAddr(req *AddrRequest) (*AddrResponse, error) {
	extAddr := req.ExtAddr

	if resp, ok, err := fetchAddrFromCache(extAddr); err != nil {
		return nil, err
	} else if ok {
		return resp, nil
	}

	net, addr, ok := util.DecodeExtId(extAddr)
	if !ok {
		return nil, &InputError{Message: "malformed ext address"}
	}
	fetcher := network.GetFetcher(net)
	if fetcher == nil {
		return nil, &InputError{Message: "not supported network"}
	}
	binaryHex, err := fetcher.GetCode(addr)
	if err != nil {
		log.Println(err)
		return nil, &NetworkError{Message: "cannot fetch contract code"}
	}
	extCodeID, err := SaveCode("evm_generic", string(binaryHex))
	if err != nil {
		return nil, err
	}

	err = storage.Set(addrTable, extAddr, []byte(extCodeID))
	return &AddrResponse{
		ExtCodeID: extCodeID,
		Binary:    string(binaryHex),
	}, err
}

func fetchAddrFromCache(extAddr string) (*AddrResponse, bool, error) {
	if value, ok, err := storage.Get(addrTable, extAddr); ok {
		extCodeID := string(value)
		if binary, err := LoadCodeID(extCodeID); err != nil {
			// Could be some race condition (saved extCodeID but no code yet)
			// Force re-fetch
			return nil, false, nil
		} else {
			return &AddrResponse{
				ExtCodeID: extCodeID,
				Binary:    hex.EncodeToString(binary),
			}, true, nil
		}
	} else {
		return nil, false, err
	}
}
