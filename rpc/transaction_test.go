package rpc

import (
	"encoding/hex"
	"fmt"
	"github.com/go-chain/go-xrp/data"
	"math/rand"
	"testing"
	"time"
)

const (
	from         = "rsTwerzJEGiKh7WjJcC3Q7776D4eGvDXPz"
	to           = "rG5AB117rJ7e2MZGKE4XfaVK5BdyHBxcSm"
	from_private = "c1dd3d6ba77aaa58a15f706e54c4d1dd59b2f61fdb621fd2dd80e204f0eaa2dc"
	to_private   = "c57891a6f2212dd312a12cb9323e69b6ad8a0faaf8435ca533876a7c12b80ae8"
)

func TestClient_Submit(t *testing.T) {
	fromAccount, _ := data.NewAccountFromAddress(to)
	toAccount, _ := data.NewAccountFromAddress(from)

	account, err := client.GetAccountInfo(to)
	if err != nil {
		panic(err)
	}
	fmt.Println("get account seq: ", account.AccountData.Sequence)

	server, err := client.GetServerInfo()
	if err != nil {
		panic(err)
	}

	amount, _ := data.NewAmount("3.9" + "/XRP")
	feeVal, _ := data.NewValue("12", true)

	last := server.State.ValidatedLedger.Seq + 100

	fmt.Printf("last ledger seq: %d add 10 after: %d\n", server.State.ValidatedLedger.Seq, last)
	rand.Seed(time.Now().Unix())
	tag := uint32(rand.Intn(100000))
	fmt.Printf("tag: %d\n", tag)

	flags := data.TransactionFlag(2147483648)
	txnBase := data.TxBase{
		TransactionType:    data.PAYMENT,
		Account:            *fromAccount,
		Sequence:           account.AccountData.Sequence,
		Fee:                *feeVal,
		LastLedgerSequence: &last,
		Flags:              &flags,
		SourceTag:          &tag,
	}

	payment := &data.Payment{
		TxBase:         txnBase,
		Destination:    *toAccount,
		Amount:         *amount,
		DestinationTag: &tag,
	}

	txBlob, err := client.signOffline(payment, to_private)
	if err != nil {
		panic(err)
	}
	fmt.Println("raw tx: ", txBlob)

	tx, err := client.Submit(txBlob)
	if err != nil {
		panic(err)
	}
	fmt.Println("txid: ", tx.TxJson.Hash)
}

func TestSign(t *testing.T) {
	//from := "rBWXYuhqESshBv6a29sMqJ59yrotwpsupf"
	//to := "rBPgesZDaeWHHtdKdcTfakxHKYhXQCgDdq"
	//value := "0.1"
	//currency := "XRP"
	//Fee := "12"
	//pri := "D61AFD08C77AC3769CD2B3A7DD44B966B092A74605FE61C0973264B5B2D53DB3"
	//
	//s, err := client.Sign(from, to, currency, value,Fee, pri,15,13313150)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//}
	//fmt.Println(s)

}

//00D61AFD08C77AC3769CD2B3A7DD44B966B092A74605FE61C0973264B5B2D53DB3
//86029426A6D950A14CEDD1AE33F0EB8C7CE1C0E8190D41D82C52EA160084B9E8

// tx_blob
// 1200002280000000240000000F201B00CB247E6140000000000186A068400000000000000C7321028C35EEA94EE7FA9C8485426E164159330BA2453368F399669D5110009F270EE974473045022100D59891D15129AFA2297506207AF14A97C2C236C690BA5E167E84BC070CA3774202203F80DFC3D8965AA4705940B9233ED8570812557F1E9DC011DEAF47DC2AE8BD588114190BA3E39BE7E7267AF0C79B2E3E2BDEC738A154831424F9C8900B8C33E55A2B848587884B10EF9992C7
func TestMakeBlob(t *testing.T) {
	Account := "rsTwerzJEGiKh7WjJcC3Q7776D4eGvDXPz"
	Amount := "30"
	Destination := "rG5AB117rJ7e2MZGKE4XfaVK5BdyHBxcSm"
	Fee := "12"
	Flags := 2147483648
	last := uint32(13313150)
	Sequence := 1
	//TxnSignature := "3045022100D59891D15129AFA2297506207AF14A97C2C236C690BA5E167E84BC070CA3774202203F80DFC3D8965AA4705940B9233ED8570812557F1E9DC011DEAF47DC2AE8BD58"
	SigningPubKey := "028C35EEA94EE7FA9C8485426E164159330BA2453368F399669D5110009F270EE9"

	fromAccount, _ := data.NewAccountFromAddress(Account)
	toAccount, _ := data.NewAccountFromAddress(Destination)
	amount, _ := data.NewAmount(Amount + "/XRP")
	fee, _ := data.NewValue(Fee, true)
	flags := data.TransactionFlag(Flags)
	//tSig, _ := hex.DecodeString(TxnSignature)
	//txnSign := data.VariableLength(tSig)
	signPubKey := data.PublicKey{}
	pk, _ := hex.DecodeString(SigningPubKey)
	copy(signPubKey[:], pk)

	txn := data.TxBase{
		TransactionType:    data.PAYMENT,
		Account:            *fromAccount,
		LastLedgerSequence: &last,
		Flags:              &flags,
		Sequence:           uint32(Sequence),
		//TxnSignature:       &txnSign,
		Fee:           *fee,
		SigningPubKey: &signPubKey,
	}
	payment := data.Payment{
		TxBase:      txn,
		Amount:      *amount,
		Destination: *toAccount,
	}

	res, err := client.makeTxBlob(&payment)
	if err != nil {
		t.Error("gen blob err: ", err)
	}
	t.Log("tx blog: ", res)
}

func TestTX(t *testing.T) {
	hash := "A0563D2067382ADC7FD8115E4E2CBE8DE9C5A08658F2C0F6B3A0099B0F59BC82"
	res, err := client.TX(hash)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestPayments(t *testing.T) {
	t1 := "2018-10-11 00:00:00"
	t2 := "2018-10-11 13:00:00"
	start, _ := time.Parse("2006-01-02 15:04:05", t1)
	end, _ := time.Parse("2006-01-02 15:04:05", t2)
	resp, err := client.Payments("XRP", start.Unix(), end.Unix(), 1, "")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", resp)
}
