package ethereum

import (
	"../utils"
)

func GenerateKey() ([]byte, error) {
	return utils.GenerateKey()
}

func GetPublicKey(key []byte) []byte {
	return utils.GetPublicKeyUncompressed(key)
}

func GetAddress(key []byte) []byte {
	pbk := GetPublicKey(key)
	return utils.Keccak256(pbk[1:])[12:32]
}
