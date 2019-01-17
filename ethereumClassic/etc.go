package ethereumClassic

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/krboktv/blockchain-swiss-knife/utils"
)

type EthereumClassic struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

func (etc *EthereumClassic) GenerateKey() ([]byte, error) {
	return utils.GenerateKeySecp256k1()
}

func (etc *EthereumClassic) GetPublicKey(key []byte) []byte {
	return utils.GetPublicKeyUncompressedSecp256k1(key)
}

func (etc *EthereumClassic) GetAddress(key []byte) []byte {
	pbk := etc.GetPublicKey(key)
	return utils.Keccak256(pbk[1:])[12:32]
}

func (etc *EthereumClassic) GenerateAndSet() {
	privateKey, err := etc.GenerateKey()
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := etc.GetPublicKey(privateKey)
	address := etc.GetAddress(privateKey)

	etc.PrivateKey = hex.EncodeToString(privateKey)
	etc.PublicKey = hex.EncodeToString(publicKey)
	etc.Address = "0x" + hex.EncodeToString(address)
}

func (etc *EthereumClassic) GetBalance(address string) (balanceFloat float64) {

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
