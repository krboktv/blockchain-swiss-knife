package stellar

import (
	"fmt"
	"strconv"

	"github.com/imroc/req"
)

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
