package bitcoin

import (
	"github.com/krboktv/blockchain-swiss-knife/utils"
	"github.com/imroc/req"
	"fmt"
	"strconv"
)

func GenerateKey() ([]byte, error) {
	return utils.GenerateKey()
}

func GetPublicKey(key []byte) []byte {
	return utils.GetPublicKey(key)
}

func GetAddress(key []byte) ([]byte, error) {
	networkByte := []byte{0x00}
	pbk := GetPublicKey(key)
	step1 := utils.SHA256(pbk)
	step2 := utils.RIPEMD160(step1)
	step3 := append(networkByte, step2...)
	step4 := utils.DoubleSHA256(step3)
	step5 := append(step3, step4[:4]...)
	return utils.EncodeToBase58(utils.EncodeBitcoin, step5)
}

func GetBalance(address string)(balanceFloat float64){
	balance, err := req.Get( "https://insight.bitpay.com/api/addr/" + address + "/balance")
	if err != nil {
		fmt.Println(err)
	}

	balanceFloat, _ = strconv.ParseFloat(balance.String(),64)

	balanceFloat *= 0.00000001 // satoshi to btc

	return
}