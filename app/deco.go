package app

import (
	"encoding/hex"
	"encoding/json"

	"github.com/blukat29/evm-explorer/deco"
	"github.com/blukat29/evm-explorer/storage"
)

// extCodeID -> decompilation
const decoTable = "deco"

func Deco(req *DecoRequest) (*DecoResponse, error) {
	extCodeID := req.ExtCodeID

	var contract Contract
	cachedDoc, ok, err := storage.Get(decoTable, extCodeID)
	if err != nil {
		return nil, err
	} else if ok {
		err := json.Unmarshal(cachedDoc, &contract)
		return &DecoResponse{Contract: contract}, err
	}

	code, err := LoadCodeID(extCodeID)
	if err != nil {
		return nil, err
	}
	codeHex := hex.EncodeToString(code)

	doc, err := deco.RunWorker(codeHex)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(doc, &contract); err != nil {
		return nil, err
	}

	err = storage.Set(decoTable, extCodeID, doc)
	return &DecoResponse{Contract: contract}, err
}
