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
	extCodeID := req.ExtCodeID

	var contract Contract
	if deco, ok := decoDB[extCodeID]; ok {
		if err := json.Unmarshal([]byte(deco.CodeJson), &contract); err != nil {
			return nil, err
		}
		return &DecoResponse{Contract: contract}, nil
	}

	binCode, ok := codeDB[extCodeID]
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

	decoDB[extCodeID] = &Decompilation{
		CodeJson: string(contractJson),
	}
	return &DecoResponse{Contract: contract}, nil
}
