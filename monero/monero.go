package monero

import (
	"encoding/hex"
	"github.com/krboktv/blockchain-swiss-knife/utils"
)

type Monero struct {
	Seed            string
	PrivateKeyVeiw  string
	PrivateKeySpend string
	PublicKeyVeiw   string
	PublicKeySpend  string
	Address         string
}

var (
	MainnetXMR = []byte{0x12}
)

func (monero *Monero) GenerateRandomSeed() string {
	return utils.GenerateRandomMnemonic(utils.DictionaryMonero,25)
}

//func (stellar *Monero) GenerateKey() ([]byte, error) {
//	return
//}

func (monero *Monero) GetPrivateKeyFromSeed(seed string) ([]byte, error) {
	seedHex, err := hex.DecodeString(seed)
	if err != nil {
		return nil, err
	}
	pvk, _, err := utils.GenerateKeyEd25519FromSeed(seedHex)
	return pvk, err
}

