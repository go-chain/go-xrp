package rpc

type LedgerCurrentResult struct {
	Reslut *LedgerCurrent `json:"result"`
}
type LedgerCurrent struct {
	LedgerCurrentIndex uint32 `json:"ledger_current_index"`
	Status string `json:"status"`
}

type LedgerRequest struct {
	LedgerIndex uint32 `json:"ledger_index"`
	Accounts bool `json:"accounts"`
	Full bool `json:"full"`
	Transactions bool `json:"transactions"`
	Expand bool `json:"expand"`
	OwnerFunds bool `json:"owner_funds"`
}

type LedgerResult struct {
	Reslut *Ledger `json:"result"`
}

type Ledger struct {
	LLedger *LedgerSub `json:"ledger"`
	LedgerHash string `json:"ledger_hash"`
	LedgerIndex uint32 `json:"ledger_index"`
	Status string `json:"status"`
	Validated bool `json:"validated"`
}

type  LedgerSub struct {
	Accepted  bool `json:"accepted"`
	AccountHash string `json:"account_hash"`
	CloseFlags uint32 `json:"close_flags"`
	CloseTime uint32 `json:"close_time"`
	CloseTimeHuman string `json:"close_time_human"`
	CloseTimeResolution uint32 `json:"close_time_resolution"`
	Closed bool `json:"closed"`
	Hash string `json:"hash"`
	LedgerHash string `json:"ledger_hash"`
	LedgerIndex string `json:"ledger_index"`
	ParentCloseTime uint32 `json:"parent_close_time"`
	ParentHash string `json:"parent_hash"`
	SeqNum string `json:"seqNum"`
	TotalCoins string `json:"totalCoins"`
	TransactionHash string `json:"transaction_hash"`
	Transactions []string `json:"transactions"`
} 
