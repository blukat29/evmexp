package app

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/blukat29/evm-explorer/util"
)

type BinaryCode struct {
	Format string
	Binary []byte // not hex; binary data.
}

var codeDB = map[string]*BinaryCode{}

func SaveCode(format, codeHex string) (string, error) {
	codeHex = strings.TrimSpace(codeHex)
	if strings.HasPrefix(codeHex, "0x") {
		codeHex = codeHex[2:]
	}
	binary, err := hex.DecodeString(codeHex)
	if err != nil {
		return "", &InputError{Message: "malformed binary"}
	}

	codeID := fmt.Sprintf("%x", sha256.Sum256(binary))
	extCodeID := util.EncodeExtId(format, codeID)

	if _, ok := codeDB[extCodeID]; !ok {
		codeDB[extCodeID] = &BinaryCode{
			Format: format,
			Binary: binary,
		}
	}

	return extCodeID, nil
}

func CodeUpload(req *CodeUploadRequest) (*CodeUploadResponse, error) {
	// Parse input
	if len(req.Format) == 0 {
		req.Format = "evm_generic"
	}
	switch req.Format {
	case "evm_generic":
	default:
		return nil, &InputError{Message: "invalid format"}
	}

	if len(req.Binary) == 0 {
		return nil, &InputError{Message: "empty binary"}
	}
	extCodeID, err := SaveCode(req.Format, req.Binary)
	if err != nil {
		return nil, err
	}

	return &CodeUploadResponse{
		ExtCodeID: extCodeID,
	}, nil
}
