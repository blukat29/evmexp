package network

import (
	"github.com/blukat29/evm-explorer/util"
	"github.com/ethereum/go-ethereum/rpc"
)

type EthFetcher struct {
	Endpoint string
}

func NewEthFetcher(endpoint string) *EthFetcher {
	return &EthFetcher{
		Endpoint: endpoint,
	}
}

func (f *EthFetcher) GetCode(addr string) ([]byte, error) {
	client, err := rpc.Dial(f.Endpoint)
	if err != nil {
		return nil, err
	}

	var result string
	err = client.Call(&result, "eth_getCode", addr, "latest")
	return []byte(util.RemoveHexPrefix(result)), err
}
