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
	format := req.Format

	if len(req.Binary) == 0 {
		return nil, &InputError{Message: "empty binary"}
	}
	binStr := strings.TrimSpace(req.Binary)
	if strings.HasPrefix(binStr, "0x") {
		binStr = binStr[2:]
	}
	binary, err := hex.DecodeString(binStr)
	if err != nil {
		return nil, &InputError{Message: "malformed binary"}
	}

	codeHash := fmt.Sprintf("%x", sha256.Sum256(binary))
	extendedCodeHash := format + "-" + codeHash

	if _, ok := codeDB[extendedCodeHash]; !ok {
		codeDB[extendedCodeHash] = &BinaryCode{
			Format: format,
			Binary: binary,
		}
	}
	return &CodeUploadResponse{
		ExtendedCodeHash: extendedCodeHash,
	}, nil
}
