package network

type Fetcher interface {
	GetCode(addr string) ([]byte, error)
}

var fetchers = map[string]Fetcher{}

func Init() error {
	// https://ethereumnodes.com/
	fetchers["eth-mainnet"] = NewEthFetcher("https://cloudflare-eth.com/")
	return nil
}

func GetFetcher(name string) Fetcher {
	return fetchers[name]
}
