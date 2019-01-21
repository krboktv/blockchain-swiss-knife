package zcash

import (
	"encoding/hex"
	"fmt"
	"github.com/imroc/req"
	"github.com/krboktv/blockchain-swiss-knife/utils"
)

type ZCash struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

var MainnetZCash = []byte{0x1C, 0xB8}

func (zcash *ZCash) GenerateKey() ([]byte, error) {
	return utils.GenerateKeySecp256k1()
}

func (zcash *ZCash) GetPublicKey(key []byte) []byte {
	return utils.GetPublicKeySecp256k1(key)
}

func (zcash *ZCash) GetAddress(key []byte) ([]byte, error) {
	pbk := zcash.GetPublicKey(key)
	step1 := utils.SHA256(pbk)
	step2 := utils.RIPEMD160(step1)
	step3 := append(MainnetZCash, step2...)
	step4 := utils.DoubleSHA256(step3)
	step5 := append(step3, step4[:4]...)
	return utils.EncodeToBase58(utils.EncodeBitcoin, step5)
}

func (zcash *ZCash) GenerateAndSet() {
	privateKey, err := zcash.GenerateKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := zcash.GetPublicKey(privateKey)

	address, err := zcash.GetAddress(privateKey)

	zcash.PrivateKey = hex.EncodeToString(privateKey)
	zcash.PublicKey = hex.EncodeToString(publicKey)
	zcash.Address = string(address)
}

func (zcash *ZCash) GetBalance(address string) (balanceFloat float64) {
	//
	type ZCashBalance struct {
		Address    string  `json:"address"`
		Balance    float64 `json:"balance"`
		FirstSeen  int     `json:"firstSeen"`
		LastSeen   int     `json:"lastSeen"`
		SentCount  int     `json:"sentCount"`
		RecvCount  int     `json:"recvCount"`
		MinedCount int     `json:"minedCount"`
		TotalSent  float64 `json:"totalSent"`
		TotalRecv  float64 `json:"totalRecv"`
	}

	var balance ZCashBalance

	resp, err := req.Get("https://api.zcha.in/v2/mainnet/accounts/" + address)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.ToJSON(&balance)

	balanceFloat = balance.Balance

	return
}
