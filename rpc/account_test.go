package rpc

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	testNet = "https://s.altnet.rippletest.net:51234"
	myTestNet = "http://47.74.211.50:5005"
)

var (
	client = NewClient(myTestNet, "https://testnet.data.api.ripple.com")
	//client = NewClient("http://47.75.70.201:9003", "http://47.75.70.201:9003")
	//client = NewClient("https://data.ripple.com")
)

func TestGetAccountBalance(t *testing.T) {
	address := "rsTwerzJEGiKh7WjJcC3Q7776D4eGvDXPz"
	res, err := client.GetAccountBalances(address, map[string]string{})
	if err != nil {
		t.Error("get err: ", err)
	}
	for _, v := range res.Balances {
		fmt.Printf("balance: %+v\n", v)
	}
}

func TestGetAccountInfo(t *testing.T) {
	address := "rsTwerzJEGiKh7WjJcC3Q7776D4eGvDXPz"
	res, err := client.GetAccountInfo(address)
	if err != nil {
		t.Error("err: ", err)
	}
	fmt.Printf("res: %+v\n", res.AccountData)
}

func TestGenAddress(t *testing.T) {
	pri, pub, addr, err := client.GenAddress()
	if err != nil {
		t.Error(err)
	}
	t.Log("pri: ", pri)
	t.Log("pub: ", pub)
	t.Log("addr: ", addr)
}

func TestGetAccountTransaction(t *testing.T) {
	address := "rBWXYuhqESshBv6a29sMqJ59yrotwpsupf"
	resp, err := client.GetAccountTransactions(address, map[string]string{"limit": "30"})
	if err != nil {
		t.Error(err)
	}
	res, _ := json.Marshal(resp)
	fmt.Printf("%s\n", res)
}

//pri:  d3c34cc4553591860f14fb64dd9562210f57b2e12970e752e54402fc7dd2844f
//pub:  0373330fcc500d6e7b1ce775ac9ca2cfa13b805befcd5b17d5108d1246d1bb6337
//addr:  raMjJMN8LDogUx6BckDV7LojR8XAXZDm91


//pri:  c57891a6f2212dd312a12cb9323e69b6ad8a0faaf8435ca533876a7c12b80ae8
//pub:  033b849512a08922e93c74d43c13c6b1c2dc8591bf787b4f36e81d511bea587dd6
//addr:  rG5AB117rJ7e2MZGKE4XfaVK5BdyHBxcSm