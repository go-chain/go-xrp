package rpc

import (
	"encoding/json"
	"go-ripple/tools/http"
)

func (c *Client) GetLedger(index uint32) (*Ledger, error) {

	params, err := json.Marshal(LedgerRequest{LedgerIndex:index,Accounts:false,Full:false,Transactions:true,Expand:false,OwnerFunds:false})
	if err != nil {
		return nil,err
	}
	//fmt.Println("params: ",string(params))
	resp, err := http.HttpPost(c.rpcJsonURL, []byte(`{"method":"ledger", "params": [`+string(params)+`]}`))
	if err != nil {
		return nil, err
	}
	res := &LedgerResult{}
	err = json.Unmarshal(resp, res)
	return res.Reslut, err
}

func (c *Client) GetLedgerCurrent() (*LedgerCurrent,error) {
	resp, err := http.HttpPost(c.rpcJsonURL, []byte(`{"method":"ledger_current", "params": [{}]}`))
	if err != nil {
		return nil, err
	}
	res := &LedgerCurrentResult{}
	err = json.Unmarshal(resp, res)
	return res.Reslut, err
}