package network

import (
	"encoding/base64"

	"github.com/ethereum/go-ethereum/rpc"
)

type KlayFetcher struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
}

func NewKlayFetcher(endpoint string) *KlayFetcher {
	return &KlayFetcher{
		Endpoint:        endpoint,
		AccessKeyId:     "KASK..",
		AccessKeySecret: "...",
	}
}

func authHeader(username, password string) string {
	concat := username + ":" + password
	encoded := base64.StdEncoding.EncodeToString([]byte(concat))
	return "Basic " + encoded
}

func (f *KlayFetcher) GetCode(addr string) ([]byte, error) {
	// https://refs.klaytnapi.com/ko/node/latest
	// https://ko.docs.klaytn.com/bapp/json-rpc/api-references/klay/account#klay_getcode
	client, err := rpc.Dial(f.Endpoint)
	if err != nil {
		return nil, err
	}

	client.SetHeader("x-chain-id", "8217")
	client.SetHeader("Authorization",
		authHeader(f.AccessKeyId, f.AccessKeySecret))

	var result string
	err = client.Call(&result, "klay_getCode", addr, "latest")
	return []byte(result), err
}
