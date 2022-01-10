package deco

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
)

var workerPath = "./worker/main.py"

func RunWorker(code string) ([]byte, error) {
	cmd := exec.Command(workerPath, "--stdin")

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	go func() {
		io.WriteString(stdin, code+"\n")
		stdin.Close()
	}()

	if err := cmd.Run(); err != nil {
		dumpOutputs(outb, errb)
		return nil, fmt.Errorf("worker error: %s", err)
	}

	var doc interface{}
	out := outb.Bytes()
	if err := json.Unmarshal(out, &doc); err != nil {
		dumpOutputs(outb, errb)
		return nil, fmt.Errorf("worker output is not json: %s", err)
	}

	return out, nil
}

func dumpOutputs(outb, errb bytes.Buffer) {
	log.Println("----- worker stdout -----\n", outb.String(), "\n----------\n")
	log.Println("----- worker stderr -----\n", errb.String(), "\n----------\n")
}
