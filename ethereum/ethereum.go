package ethereum

import (
	"github.com/krboktv/blockchain-swiss-knife/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"math/big"
	"log"
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

func GetBalance(addr string)(balance *big.Int){
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	account := common.HexToAddress(addr)
	balance, err = client.BalanceAt(ctx, account, nil)
	if err != nil {
		log.Fatal(err)
	}
	return
}
