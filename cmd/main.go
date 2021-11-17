package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/blukat29/evm-explorer/app"
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
		res, err := app.Decompile(&app.ContractRequest{
			Addr:  "",
			Chain: "",
			Code:  string(code),
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Contract)
		return
	}

	app.Serve()
}
