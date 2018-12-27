package utils

import (
	"crypto"
	"crypto/rand"
	eth "github.com/ethereum/go-ethereum/crypto"
	"github.com/haltingstate/secp256k1-go"
	"github.com/spearson78/guardian/encoding/base58"
	"github.com/vsergeev/btckeygenie/btckey"
)

const EncodeRipple = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
const EncodeBitcoin = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

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

func EncodeToBase58(alphabet string, address []byte) ([]byte, error) {
	return base58.NewEncoding(alphabet).Encode(address)
}
