package deco

import (
	"io"
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

	return cmd.Output()
}
