package rpc

import (
	"encoding/json"
	"github.com/go-chain/go-xrp/tools/http"
)

func (c *Client) GetServerInfo() (*ServerInfoResult, error) {
	resp, err := http.HttpPost(c.rpcJsonURL, []byte(`{"method":"server_state", "params": [{}]}`))
	if err != nil {
		return nil, err
	}
	res := &ServerInfoResp{}
	err = json.Unmarshal(resp, res)
	return res.Result, err
}
