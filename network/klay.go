package network

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/blukat29/evm-explorer/util"
	"github.com/ethereum/go-ethereum/rpc"
)

type KlayFetcher struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
}

type KlayFetcherOptions struct {
	AccessKeyId     string
	AccessKeySecret string
}

func NewKlayFetcher(endpoint string, options *KlayFetcherOptions) *KlayFetcher {
	f := &KlayFetcher{Endpoint: endpoint}

	if options != nil {
		f.AccessKeyId = options.AccessKeyId
		f.AccessKeySecret = options.AccessKeySecret
	}

	if len(f.AccessKeyId) == 0 || len(f.AccessKeySecret) == 0 {
		log.Fatal("No KAS API key specified")
	}
	fmt.Println(f.AccessKeyId, f.AccessKeySecret)

	return f
}

func authHeader(username, password string) string {
	concat := username + ":" + password
	encoded := base64.StdEncoding.EncodeToString([]byte(concat))
	fmt.Println(encoded)
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
	return []byte(util.RemoveHexPrefix(result)), err
}
