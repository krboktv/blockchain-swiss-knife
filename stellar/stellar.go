package stellar

import (
	"encoding/binary"
	"fmt"
	"github.com/krboktv/blockchain-swiss-knife/utils"
	"strconv"
	"crypto/rand"
	"github.com/imroc/req"
)

func GenerateKey() ([]byte, error) {
	return generateRandomSeed()
}

func generateRandomPassphrase() ([]byte, error) {
	phrase := make([]byte, 32)
	rnd, err := rand.Read(phrase)
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint32(phrase, uint32(rnd))
	return phrase, nil
}

func generateRandomSeed() ([]byte, error) {
	passphrase, err := generateRandomPassphrase()
	if err != nil {
		return nil, err
	}
	seed := utils.SHA512(passphrase)[:32]
	return encodeSeedToBase58Check(seed)
}

func generateSeedFromPassphrase(passphrase []byte) ([]byte, error) {
	seed := utils.SHA512(passphrase)[:32]
	return encodeSeedToBase58Check(seed)
}

func encodeSeedToBase58Check(seed []byte) ([]byte, error) {
	networkByte := []byte{0x33}
	step1 := append(networkByte, seed...)
	step2 := utils.DoubleSHA256(step1)
	step3 := append(step1, step2[:4]...)
	step4, err := utils.EncodeToBase58(utils.EncodeStellar, step3)
	return step4, err
}

func GetBalance(address string) (balanceFloat float64) {

	type StellarBalance struct {
		Balances []struct {
			Balance             string `json:"balance"`
			Buying_liabilities  string `json:"buying_liabilities"`
			Selling_liabilities string `json:"selling_liabilities"`
			Asset_type          string `json:"asset_type"`
		}
	}

	var b StellarBalance
	balance, err := req.Get("https://horizon.stellar.org/accounts/" + address)
	if err != nil {
		fmt.Println(err)
	}

	balance.ToJSON(&b)

	var stellarBalanceString string

	for _, j := range b.Balances {
		if j.Asset_type == "native" {

			stellarBalanceString = j.Balance
		}
	}

	balanceFloat, err = strconv.ParseFloat(stellarBalanceString, 64)
	if err != nil {
		fmt.Println(err)
	}

	return
}
