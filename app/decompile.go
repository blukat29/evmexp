package app

import (
	"encoding/json"

	"github.com/blukat29/evm-explorer/deco"
)

func Decompile(req *ContractRequest) (*ContractResponse, error) {
	contractJson, err := deco.RunWorker(req.Code)
	if err != nil {
		return nil, err
	}

	var contract Contract
	if err := json.Unmarshal(contractJson, &contract); err != nil {
		return nil, err
	}

	return &ContractResponse{
		Contract: contract,
	}, nil
}
