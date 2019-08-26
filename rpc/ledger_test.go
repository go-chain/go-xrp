package rpc

import (
	"fmt"
	"testing"
	"time"
)

func TestClient_GetLedgerCurrent(t *testing.T) {
	r, err := client.GetLedgerCurrent()
	if err != nil {
		panic(err)
	}
	fmt.Println(r.LedgerCurrentIndex)
}

func TestClient_GetLedger(t *testing.T) {

	var count uint32 = 15186011
	for {
		<- time.After(2*time.Second)
		fmt.Printf("get ledger index: %d\n",count)
		resp, err := client.GetLedger(count)
		if err != nil {
			panic(err)
		}
		//fmt.Println("r: ",resp.LLedger.Transactions)
		if len(resp.LLedger.Transactions) == 0 {
			count++
			continue
		}
		for _, v := range resp.LLedger.Transactions {
			r, err := client.TX(v)
			if err != nil {
				panic(err)
			}
			fmt.Printf("ledger_index: %d from: %s to: %s money: %s txid: %s tag: %d\n",r.LedgerIndex, r.Account, r.Destination, r.Amount, r.Hash, r.DestinationTag)
		}
		count++
	}

}

func TestClient_TX(t *testing.T) {
	r, err := client.TX("144F1AE18E269A7D871C46DDF0E28B3E929B449C40AA437EC376A3675F0E3F6B")
	if err != nil {
		panic(err)
	}
	fmt.Printf("from: %s to: %s money: %s txid: %s tag: %d\n",r.Account,r.Destination,r.Amount,r.Hash,r.DestinationTag)
}