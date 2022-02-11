package model

type Transaction struct {
	Hash string `json:"hash"`
	Nonce string `json:"nonce"`
	BlockHash string `json:"blockHash"`
	BlockNumber string `json:"blockNumber"`
	TransactionIndex string `json:"transactionIndex"`
	From string `json:"from"`
	To string `json:"to"`
	Value string `json:"value"`
	Gasprice string `json:"gasprice"`
	Gas string `json:"gas"`
	Input string `json:"input"`
	Difficulty string `json:"difficulty"`
	TotalDifficuly string `json:"totalDifficuly"`
	ExtraData string `json:"extraData"`
	Size string `json:"size"`
	GasLimit string `json:"gasLimit"`
	GasUsed string `json:"gasUsed"`
	Timestamp string `json:"timestamp"`
	Uncles []string `json:"uncles"`
	Transactions []interface{} `json:"transactions"`
}

