package app

import (
	"encoding/hex"
	"encoding/json"

	"github.com/blukat29/evm-explorer/deco"
)

type Decompilation struct {
	CodeJson string
}

var decoDB = map[string]*Decompilation{}

func Deco(req *DecoRequest) (*DecoResponse, error) {
	extendedCodeHash := req.ExtendedCodeHash

	var contract Contract
	if deco, ok := decoDB[extendedCodeHash]; ok {
		if err := json.Unmarshal([]byte(deco.CodeJson), &contract); err != nil {
			return nil, err
		}
		return &DecoResponse{Contract: contract}, nil
	}

	binCode, ok := codeDB[extendedCodeHash]
	if !ok {
		return nil, &NotFoundError{Message: "no such code"}
	}
	codeHex := hex.EncodeToString(binCode.Binary)

	contractJson, err := deco.RunWorker(codeHex)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(contractJson, &contract); err != nil {
		return nil, err
	}

	decoDB[extendedCodeHash] = &Decompilation{
		CodeJson: string(contractJson),
	}
	return &DecoResponse{Contract: contract}, nil
}
