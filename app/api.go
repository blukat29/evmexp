package app

// In-memory types

type InputError struct {
	Message string
}

func (e *InputError) Error() string {
	return e.Message
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type NetworkError struct {
	Message string
}

func (e *NetworkError) Error() string {
	return e.Message
}

// JSON data types

type Response struct {
	Error string `json:"error"`
}

type AddrRequest struct {
	ExtAddr string `uri:"addr" binding:"required"`
}

type AddrResponse struct {
	Response
	ExtCodeID string `json:"extCodeID"`
	Binary    string `json:"binary"`
}

type CodeUploadRequest struct {
	Format string `json:"format"`
	Binary string `json:"binary" binding:"required"`
}

type CodeUploadResponse struct {
	Response
	ExtCodeID string `json:"extCodeID"`
}

// TODO: further concretize Functions and Storages.
type Contract struct {
	Asm        string        `json:"asm"`
	Pseudocode string        `json:"pseudocode"`
	Functions  []interface{} `json:"functions"`
	Storages   []interface{} `json:"storage"`
}

type DecoRequest struct {
	ExtCodeID string `uri:"id" binding:"required"`
}

type DecoResponse struct {
	Response
	Contract interface{} `json:"contract"`
}
