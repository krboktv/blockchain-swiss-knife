package ethereumClassic

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/krboktv/blockchain-swiss-knife/utils"
)

func GenerateKey() ([]byte, error) {
	return utils.GenerateKeySecp256k1()
}

func GetPublicKey(key []byte) []byte {
	return utils.GetPublicKeyUncompressedSecp256k1(key)
}

func GetAddress(key []byte) []byte {
	pbk := GetPublicKey(key)
	return utils.Keccak256(pbk[1:])[12:32]
}

func GetBalance(address string) (balanceFloat float64) {

	client, err := ethclient.Dial("https://ethereumclassic.network")
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()

	account := common.HexToAddress(address)

	balance, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		fmt.Println(err)
	}
	ethBalance, _ := strconv.ParseFloat(balance.String(), 64)

	balanceFloat = ethBalance / math.Pow(10, 18)

	return
}
