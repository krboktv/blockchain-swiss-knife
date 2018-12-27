package stellar

import (
	"github.com/sirupsen/logrus"
	"github.com/stellar/go/keypair"
)

type KeyPair struct {
	Seed    string // private key
	Address string // public key
}

func debugf(method string, msg string, args ...interface{}) {
	logrus.WithFields(logrus.Fields{"lib": "microstellar", "method": method}).Debugf(msg, args...)
}

func GeneratePrivateKey()(*KeyPair, error){
	pair, err := keypair.Random()
	if err != nil {
		return nil,err
	}

	debugf("CreateKeyPair", "created address: %s, seed: <redacted>", pair.Address())
	return &KeyPair{pair.Seed(), pair.Address()}, nil
}