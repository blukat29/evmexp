package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/blukat29/evm-explorer/deco"
)

func main() {
	path := os.Args[1]
	code, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	d, err := deco.RunWorker(string(code))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(d))
}
