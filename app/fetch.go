package app

import (
	"strings"

	"github.com/blukat29/evm-explorer/network"
)

type AddrInfo struct {
	ExtendedCodeHash string
}

var addrDB = map[string]*AddrInfo{}

func FetchAddr(req *AddrRequest) (*AddrResponse, error) {
	extendedAddr := req.ExtendedAddr

	if info, ok := addrDB[extendedAddr]; ok {
		return &AddrResponse{
			ExtendedCodeHash: info.ExtendedCodeHash,
		}, nil
	}

	parts := strings.Split(extendedAddr, "-")
	if len(parts) != 3 {
		return nil, &InputError{Message: "malformed extended address"}
	}
	fetcher := network.GetFetcher(parts[0] + "-" + parts[1])
	if fetcher == nil {
		return nil, &InputError{Message: "not supported network"}
	}
	code, err := fetcher.GetCode(parts[2])
	if err != nil {
		return nil, err
	}
	extendedCodeHash, err := SaveCode("evm-generic", string(code))
	if err != nil {
		return nil, err
	}

	addrDB[extendedAddr] = &AddrInfo{
		ExtendedCodeHash: extendedCodeHash,
	}
	return &AddrResponse{
		ExtendedCodeHash: extendedCodeHash,
	}, nil
}
