package app

import (
	"log"

	"github.com/blukat29/evm-explorer/network"
	"github.com/blukat29/evm-explorer/storage"
	"github.com/blukat29/evm-explorer/util"
)

// extAddr -> extCodeID
const addrTable = "addr"

func FetchAddr(req *AddrRequest) (*AddrResponse, error) {
	extAddr := req.ExtAddr

	if value, ok, err := storage.Get(addrTable, extAddr); err != nil {
		return nil, err
	} else if ok {
		return &AddrResponse{ExtCodeID: string(value)}, nil
	}

	net, addr, ok := util.DecodeExtId(extAddr)
	if !ok {
		return nil, &InputError{Message: "malformed ext address"}
	}
	fetcher := network.GetFetcher(net)
	if fetcher == nil {
		return nil, &InputError{Message: "not supported network"}
	}
	code, err := fetcher.GetCode(addr)
	if err != nil {
		log.Println(err)
		return nil, &NetworkError{Message: "cannot fetch contract code"}
	}
	extCodeID, err := SaveCode("evm_generic", string(code))
	if err != nil {
		return nil, err
	}

	err = storage.Set(addrTable, extAddr, []byte(extCodeID))
	return &AddrResponse{ExtCodeID: extCodeID}, err
}
