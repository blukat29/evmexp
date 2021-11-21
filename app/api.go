package app

// In-memory types

type InputError struct {
	Message string
}

func (e *InputError) Error() string {
	return e.Message
}

// JSON data types

type Response struct {
	Error string `json:"error"`
}

// TODO: further concretize Functions and Storages.
type Contract struct {
	Asm        string        `json:"asm"`
	Pseudocode string        `json:"pseudocode"`
	Functions  []interface{} `json:"functions"`
	Storages   []interface{} `json:"storage"`
}

type ContractRequest struct {
	Addr  string `json:"addr"`
	Chain string `json:"chain"`
	Code  string `json:"code"`
}

type ContractResponse struct {
	Contract Contract `json:"contract"`
}

type CodeUploadRequest struct {
	Format string `json:"format"`
	Binary string `json:"binary" binding:"required"`
}

type CodeUploadResponse struct {
	Response
	ExtendedCodeHash string `json:"extendedCodeHash"`
}
