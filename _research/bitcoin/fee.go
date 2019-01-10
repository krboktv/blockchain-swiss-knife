package main

import (
	"github.com/imroc/req"
	"fmt"
)

type BTCFee struct {
	FastestFee int `json:"fastestFee"`
	HalfHourFee int `json:"halfHourFee"`
	HourFee int `json:"hourFee"`
}

func TxFee() (feeObj BTCFee) {
	fee, err := req.Get("https://bitcoinfees.earn.com/api/v1/fees/recommended")
	if err != nil {
		fmt.Println(err)
	}

	fee.ToJSON(&feeObj)
	return
}

func main(){
	fee := TxFee()
	fmt.Printf("fastestFee: %d\n",fee.FastestFee)
	fmt.Printf("halfHourFee: %d\n",fee.HalfHourFee)
	fmt.Printf("hourFee: %d\n",fee.HourFee)
}

