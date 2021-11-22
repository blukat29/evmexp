package app

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
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

	codeHash := fmt.Sprintf("%x", sha256.Sum256(binary))
	extendedCodeHash := format + "-" + codeHash

	if _, ok := codeDB[extendedCodeHash]; !ok {
		codeDB[extendedCodeHash] = &BinaryCode{
			Format: format,
			Binary: binary,
		}
	}

	return extendedCodeHash, nil
}

func CodeUpload(req *CodeUploadRequest) (*CodeUploadResponse, error) {
	// Parse input
	if len(req.Format) == 0 {
		req.Format = "evm-generic"
	}
	switch req.Format {
	case "evm-generic":
	default:
		return nil, &InputError{Message: "invalid format"}
	}

	if len(req.Binary) == 0 {
		return nil, &InputError{Message: "empty binary"}
	}
	extendedCodeHash, err := SaveCode(req.Format, req.Binary)
	if err != nil {
		return nil, err
	}

	return &CodeUploadResponse{
		ExtendedCodeHash: extendedCodeHash,
	}, nil
}
