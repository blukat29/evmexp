package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/blukat29/evm-explorer/app"
	"github.com/blukat29/evm-explorer/deco"
	"github.com/spf13/pflag"
)

var onlyOnce bool
var filePath string

func main() {
	pflag.BoolVar(&onlyOnce, "once", false, "Decompile once and exit")
	pflag.StringVarP(&filePath, "file", "f", "", "Bytecode file")
	pflag.Parse()

	if onlyOnce {
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

	app.Serve()
}
