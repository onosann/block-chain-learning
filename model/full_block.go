package model

type FullBLock struct {
	Number string `json:"number"`
	Hash string `json:"hash"`
	ParentHash string `json:"parentHash"`
	Nonce string `json:"nonce"`
	Sha3Uncles string `json:"sha3Uncles"`
	LogsBloom string `json:"logsbloom"`
	TransactionRoot string `json:"transactionroot"`
	ReceiptsRoot string `json:"stateRoot"`
	Miner string `json:"miner"`
}


