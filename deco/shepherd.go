package deco

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

var workerPath = "./worker/main.py"

func RunWorker(code string) ([]byte, error) {
	cmd := exec.Command(workerPath, "--stdin")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	go func() {
		io.WriteString(stdin, code+"\n")
		stdin.Close()
	}()

	if out, err := cmd.CombinedOutput(); err != nil {
		log.Println("----- worker output -----\n", string(out), "\n----------\n")
		return nil, fmt.Errorf("worker error: %s", err)
	} else {
		return out, nil
	}

}
