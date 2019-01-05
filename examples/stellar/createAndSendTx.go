package main

import (
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"fmt"
)

// Needs to be tested
func CreateTransaction(from, to, amount string) (txeB64 string) {

	tx, err := build.Transaction(
		build.SourceAccount{AddressOrSeed: from},
		build.PublicNetwork,
		//b.AutoSequence{SequenceProvider: horizon.DefaultPublicNetClient}, ???
		build.Payment(
			build.Destination{AddressOrSeed: to},
			build.NativeAmount{Amount: amount},
		),
	)
	if err != nil {
		panic(err)
	}

	txe, err := tx.Sign(from)
	if err != nil {
		panic(err)
	}

	txeB64, err = txe.Base64()
	if err != nil {
		panic(err)
	}

	return
}

func SendRawTx(blob string) int32 {

	resp, err := horizon.DefaultPublicNetClient.SubmitTransaction(blob)
	if err != nil {
		panic(err)
	}
	return resp.Ledger
}

func main(){
	tx := CreateTransaction("SCRUYGFG76UPX3EIUWGPIQPQDPD24XPR3RII5BD53DYPKZJGG43FL5HI", "GA3A7AD7ZR4PIYW6A52SP6IK7UISESICPMMZVJGNUTVIZ5OUYOPBTK6X", "0.1")
	fmt.Println(tx)
}