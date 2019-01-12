package zcash

import (
	"github.com/krboktv/blockchain-swiss-knife/utils"
)

func GenerateKey() ([]byte, error) {
	return utils.GenerateKeySecp256k1()
}

func GetPublicKey(key []byte) []byte {
	return utils.GetPublicKeySecp256k1(key)
}

func GetAddress(key []byte) ([]byte, error) {
	networkByte := []byte{0x1C,0xB8}
	pbk := GetPublicKey(key)
	step1 := utils.SHA256(pbk)
	step2 := utils.RIPEMD160(step1)
	step3 := append(networkByte, step2...)
	step4 := utils.DoubleSHA256(step3)
	step5 := append(step3, step4[:4]...)
	return utils.EncodeToBase58(utils.EncodeBitcoin, step5)
}
