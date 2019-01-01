package monero

import (
	"github.com/paxosglobal/moneroutil"
)

func GenerateKey() (prvtKey *moneroutil.Key) {
	prvtKey = moneroutil.RandomScalar()
	return
}

func GetPublicKey(key moneroutil.Key) (pubKey *moneroutil.Key) {
	pubKey = key.PubKey()
	return
}
