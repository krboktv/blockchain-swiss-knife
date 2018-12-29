package utils

import (
	"crypto"
	"crypto/rand"
	eth "github.com/ethereum/go-ethereum/crypto"
	"github.com/haltingstate/secp256k1-go"
	"github.com/spearson78/guardian/encoding/base58"
	"github.com/vsergeev/btckeygenie/btckey"
	"math/big"
)

const EncodeRipple = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
const EncodeBitcoin = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const CurveOrder = "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141"

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

func DoubleSHA256(data []byte) []byte {
	return SHA256(SHA256(data))
}

func SHA512(data []byte) []byte {
	s512 := crypto.SHA512.New()
	s512.Write(data)
	return s512.Sum(nil)
}

func Keccak256(data []byte) []byte {
	return eth.Keccak256(data)
}

func RIPEMD160(data []byte) []byte {
	r160 := crypto.RIPEMD160.New()
	r160.Write(data)
	return r160.Sum(nil)
}

func EncodeToBase58(alphabet string, data []byte) ([]byte, error) {
	return base58.NewEncoding(alphabet).Encode(data)
}

func DecodeFromBase58(alphabet string, data []byte) ([]byte, error) {
	return base58.NewEncoding(alphabet).Decode(data)
}

func AddBytes(data1 []byte, data2 []byte) (*big.Int) {
	return new(big.Int).Add(BytesToBigInt(data1), BytesToBigInt(data2))
}

func ModBytes(data1 []byte, data2 []byte) (*big.Int) {
	return new(big.Int).Mod(BytesToBigInt(data1), BytesToBigInt(data2))
}

func BytesToBigInt(data []byte) (*big.Int) {
	return new(big.Int).SetBytes(data)
}

//func IsGrater(data1 []byte, data2 []byte) (*big.Int) {
//	bool := new(big.Int).
//	return new(big.Int).SetBytes(data)
//}