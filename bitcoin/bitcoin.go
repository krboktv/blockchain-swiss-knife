package bitcoin

import (
	"encoding/hex"
	"fmt"
	"github.com/imroc/req"
	"github.com/krboktv/blockchain-swiss-knife/utils"
	"strconv"
)

var (
	MainnetBTC = []byte{0x00}
)

type Bitcoin struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

func (btc *Bitcoin) GenerateKey() ([]byte, error) {
	return utils.GenerateKeySecp256k1()
}

func (btc *Bitcoin) GetPublicKey(key []byte) []byte {
	return utils.GetPublicKeySecp256k1(key)
}

func (btc *Bitcoin) GetAddress(key []byte) ([]byte, error) {
	pbk := btc.GetPublicKey(key)
	step1 := utils.SHA256(pbk)
	step2 := utils.RIPEMD160(step1)
	step3 := append(MainnetBTC, step2...)
	step4 := utils.DoubleSHA256(step3)
	step5 := append(step3, step4[:4]...)
	return utils.EncodeToBase58(utils.EncodeBitcoin, step5)
}

func (btc *Bitcoin) GenerateAndSet() {

	privateKey, err := btc.GenerateKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := btc.GetPublicKey(privateKey)

	address, err := btc.GetAddress(privateKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	btc.PrivateKey = hex.EncodeToString(privateKey)
	btc.PublicKey = hex.EncodeToString(publicKey)
	btc.Address = string(address)

}

func (btc *Bitcoin) GetBalance(address string) (balanceFloat float64) {
	balance, err := req.Get("https://insight.bitpay.com/api/addr/" + address + "/balance")
	if err != nil {
		fmt.Println(err)
		return
	}

	balanceFloat, _ = strconv.ParseFloat(balance.String(), 64)

	balanceFloat *= 0.00000001 // satoshi to btc

	return
}
