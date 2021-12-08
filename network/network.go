package network

type Fetcher interface {
	GetCode(addr string) ([]byte, error)
}

var fetchers = map[string]Fetcher{}

func Init() error {
	// https://ethereumnodes.com/
	fetchers["eth"] = NewEthFetcher("https://cloudflare-eth.com/")
	fetchers["eth_ropsten"] = NewEthFetcher("https://ropsten-rpc.linkpool.io/")
	return nil
}

func GetFetcher(name string) Fetcher {
	return fetchers[name]
}
