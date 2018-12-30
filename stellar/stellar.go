package stellar

import (
	"github.com/krboktv/blockchain-swiss-knife/stellar/keypair"
	"github.com/sirupsen/logrus"
	"github.com/imroc/req"
	"fmt"
	"strconv"
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

func GetBalance(address string)(balanceFloat float64){

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

	var stellarBalance string

	for _, j := range b.Balances{
		if j.Asset_type == "native"{

			stellarBalance = j.Balance
		}
	}

	balanceFloat,err = strconv.ParseFloat(stellarBalance,64)
	if err != nil{
		fmt.Println(err)
	}

	return
}
