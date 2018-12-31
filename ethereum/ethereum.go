package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/krboktv/blockchain-swiss-knife/utils"
	"log"
	"math"
	"strconv"
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

func GetBalance(address string) (balanceFloat float64) {

	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	account := common.HexToAddress(address)

	balance, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		log.Fatal(err)
	}
	ethBalance, _ := strconv.ParseFloat(balance.String(), 64)

	balanceFloat = ethBalance / math.Pow(10, 18)

	return
}
