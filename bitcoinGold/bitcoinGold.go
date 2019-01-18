package bitcoinGold

import (
	"encoding/hex"
	"fmt"
	"github.com/imroc/req"
	"github.com/krboktv/blockchain-swiss-knife/utils"
	"strconv"
)

var (
	MainnetBTG = []byte{0x26}
)

type BitcoinGold struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

func (btg *BitcoinGold) GenerateKey() ([]byte, error) {
	return utils.GenerateKeySecp256k1()
}

func (btg *BitcoinGold) GetPublicKey(key []byte) []byte {
	return utils.GetPublicKeySecp256k1(key)
}

func (btg *BitcoinGold) GetAddress(key []byte) ([]byte, error) {
	pbk := btg.GetPublicKey(key)
	step1 := utils.SHA256(pbk)
	step2 := utils.RIPEMD160(step1)
	step3 := append(MainnetBTG, step2...)
	step4 := utils.DoubleSHA256(step3)
	step5 := append(step3, step4[:4]...)
	return utils.EncodeToBase58(utils.EncodeBitcoin, step5)
}

func (btg *BitcoinGold) GenerateAndSet() {
	privateKey, err := btg.GenerateKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := btg.GetPublicKey(privateKey)

	address, err := btg.GetAddress(privateKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	btg.PrivateKey = hex.EncodeToString(privateKey)
	btg.PublicKey = hex.EncodeToString(publicKey)
	btg.Address = string(address)

}

func (btg *BitcoinGold) GetBalance(address string) (balanceFloat float64) {
	balance, err := req.Get("https://explorer.bitcoingold.org/insight-api/addr/" + address + "/balance")
	if err != nil {
		fmt.Println(err)
		return
	}

	balanceFloat, _ = strconv.ParseFloat(balance.String(), 64)

	balanceFloat *= 0.00000001

	return
}
