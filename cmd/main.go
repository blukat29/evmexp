package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/blukat29/evm-explorer/app"
	"github.com/blukat29/evm-explorer/deco"
	"github.com/blukat29/evm-explorer/network"
	"github.com/blukat29/evm-explorer/storage"
	"github.com/spf13/pflag"
)

var decoOnce bool
var filePath string
var addrOnce bool
var address string

func main() {
	pflag.BoolVar(&decoOnce, "deco", false, "Decompile once and exit")
	pflag.StringVarP(&filePath, "file", "f", "", "Bytecode file")

	pflag.BoolVar(&addrOnce, "code", false, "Fetch code by address and exit")
	pflag.StringVarP(&address, "addr", "a", "", "Extended address")

	pflag.Parse()

	if decoOnce {
		code, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}
		j, err := deco.RunWorker(string(code))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(j)
		return
	}

	if addrOnce {
		network.Init()

		parts := strings.Split(address, "-")
		if len(parts) != 3 {
			log.Fatal("malformed extended address")
		}
		fetcher := network.GetFetcher(parts[0] + "-" + parts[1])
		if fetcher == nil {
			log.Fatal("not supported network")
		}
		code, err := fetcher.GetCode(parts[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(code))
		return
	}

	if err := network.Init(); err != nil {
		log.Fatal(err)
	}
	if err := storage.Init(); err != nil {
		log.Fatal(err)
	}

	app.Serve()
}
