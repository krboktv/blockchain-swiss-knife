package bitcoin

import (
	"github.com/krboktv/blockchain-swiss-knife/utils"
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
	step5 := step4[:4]
	step6 := append(step3, step5...)
	return utils.EncodeToBase58(utils.EncodeBitcoin, step6)
}
