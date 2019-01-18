package dash

import (
	"encoding/hex"
	"fmt"
	"github.com/imroc/req"
	"github.com/krboktv/blockchain-swiss-knife/utils"
	"strconv"
)

var (
	MainnetDash = []byte{0x4c}
)

type Dash struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

func (dash *Dash) GenerateKey() ([]byte, error) {
	return utils.GenerateKeySecp256k1()
}

func (dash *Dash) GetPublicKey(key []byte) []byte {
	return utils.GetPublicKeySecp256k1(key)
}

func (dash *Dash) GetAddress(key []byte) ([]byte, error) {
	pbk := dash.GetPublicKey(key)
	step1 := utils.SHA256(pbk)
	step2 := utils.RIPEMD160(step1)
	step3 := append(MainnetDash, step2...)
	step4 := utils.DoubleSHA256(step3)
	step5 := append(step3, step4[:4]...)
	return utils.EncodeToBase58(utils.EncodeBitcoin, step5)
}

func (dash *Dash) GenerateAndSet() {
	privateKey, err := dash.GenerateKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := dash.GetPublicKey(privateKey)

	address, err := dash.GetAddress(privateKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	dash.PrivateKey = hex.EncodeToString(privateKey)
	dash.PublicKey = hex.EncodeToString(publicKey)
	dash.Address = string(address)

}

func (dash *Dash) GetBalance(address string) (balanceFloat float64) {
	balance, err := req.Get("https://insight.dash.org/insight-api/addr/" + address + "/balance")
	if err != nil {
		fmt.Println(err)
		return
	}

	balanceFloat, _ = strconv.ParseFloat(balance.String(), 64)

	balanceFloat *= 0.00000001

	return
}
