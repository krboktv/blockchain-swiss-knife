package stellar

import (
	"fmt"
	"strconv"

	"github.com/imroc/req"
	"github.com/stellar/go/keypair"
	"github.com/sirupsen/logrus"
	"github.com/stellar/go/build"
)

type KeyPair struct {
	Seed    string // private key
	Address string // public key
}

func debugf(method string, msg string, args ...interface{}) {
	logrus.WithFields(logrus.Fields{"lib": "microstellar", "method": method}).Debugf(msg, args...)
}

func GenerateKey() (*KeyPair, error) {
	pair, err := keypair.Random()
	if err != nil {
		return nil, err
	}

	debugf("CreateKeyPair", "created address: %s, seed: <redacted>", pair.Address())
	return &KeyPair{pair.Seed(), pair.Address()}, nil
}

func GetBalance(address string) (balanceFloat float64) {

	type StellarBalance struct {
		Balances []struct {
			Balance             string `json:"balance"`
			Buying_liabilities  string `json:"buying_liabilities"`
			Selling_liabilities string `json:"selling_liabilities"`
			Asset_type          string `json:"asset_type"`
		}
	}

	var b StellarBalance
	balance, err := req.Get("https://horizon.stellar.org/accounts/" + address)
	if err != nil {
		fmt.Println(err)
	}

	balance.ToJSON(&b)

	var stellarBalanceString string

	for _, j := range b.Balances {
		if j.Asset_type == "native" {

			stellarBalanceString = j.Balance
		}
	}

	balanceFloat, err = strconv.ParseFloat(stellarBalanceString, 64)
	if err != nil {
		fmt.Println(err)
	}

	return
}

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
