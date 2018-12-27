package dash

import (
	"../utils"
)

func GenerateKey() ([]byte, error) {
	return utils.GenerateKey()
}

func GetPublicKey(key []byte) []byte {
	return utils.GetPublicKey(key)
}

func GetAddress(key []byte) []byte {
	networkByte := []byte{0x4c}
	pbk := GetPublicKey(key)
	step1 := utils.SHA256(pbk)
	step2 := utils.RIPEMD160(step1)
	step3 := append(networkByte, step2...)
	step4 := utils.SHA256(step3)
	step5 := utils.SHA256(step4)
	step6 := step5[:4]
	step7 := append(step3, step6...)
	return step7
}
