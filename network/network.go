package network

import (
	"log"
	"os"
)

type Fetcher interface {
	// Returns hex-encoded raw binary code.
	// In Ethereum-like networks, leading '0x' must be deleted.
	GetCode(addr string) ([]byte, error)
}

var fetchers = map[string]Fetcher{}

func Init() error {
	// https://ethereumnodes.com/
	fetchers["eth"] = NewEthFetcher("https://cloudflare-eth.com/")
	fetchers["eth_ropsten"] = NewEthFetcher("https://ropsten-rpc.linkpool.io/")

	// https://refs.klaytnapi.com/ko/node/latest
	klayAccessKeyId := os.Getenv("KAS_KEY_ID")
	klayAccessKeySecret := os.Getenv("KAS_SECRET")
	if len(klayAccessKeyId) == 0 || len(klayAccessKeySecret) == 0 {
		log.Println("Envs KAS_KEY_ID or KAS_SECRET are missing. Disabled Klaytn fetcher.")
	} else {
		fetchers["klay"] = NewKlayFetcher("https://node-api.klaytnapi.com/v1/klaytn",
			&KlayFetcherOptions{
				AccessKeyId:     klayAccessKeyId,
				AccessKeySecret: klayAccessKeySecret,
			})
	}

	return nil
}

func GetFetcher(name string) Fetcher {
	return fetchers[name]
}
