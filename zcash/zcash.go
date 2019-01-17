package zcash

import (
	"encoding/hex"
	"fmt"
	"github.com/krboktv/blockchain-swiss-knife/utils"
)

var (
	MainnetZCash = []byte{0x1C, 0xB8}
)

type ZCash struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

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
