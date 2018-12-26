package utils

import (
	"crypto"
	"crypto/rand"
	eth "github.com/ethereum/go-ethereum/crypto"
	"github.com/haltingstate/secp256k1-go"
	"github.com/vsergeev/btckeygenie/btckey"
)

func GenerateKey() ([]byte, error) {
	pvk, err := btckey.GenerateKey(rand.Reader)
	return pvk.D.Bytes(), err
}

func GetPublicKeyUncompressed(key []byte) []byte {
	return secp256k1.UncompressedPubkeyFromSeckey(key)
}

func GetPublicKey(key []byte) []byte {
	return secp256k1.PubkeyFromSeckey(key)
}

func SHA256(data []byte) []byte {
	s256 := crypto.SHA256.New()
	s256.Write(data)
	return s256.Sum(nil)
}

func Keccak256(data []byte) []byte {
	return eth.Keccak256(data)
}

func RIPEMD160(data []byte) []byte {
	r160 := crypto.RIPEMD160.New()
	r160.Write(data)
	return r160.Sum(nil)
}
