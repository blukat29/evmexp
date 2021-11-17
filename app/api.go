package app

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
