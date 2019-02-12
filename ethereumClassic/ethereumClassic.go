package ethereumClassic

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/krboktv/blockchain-swiss-knife/utils"
	"math/big"
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
		return
	}

	ctx := context.Background()

	account := common.HexToAddress(address)

	balance, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	ethBalance, _ := strconv.ParseFloat(balance.String(), 64)

	balanceFloat = ethBalance / math.Pow(10, 18)

	return
}

func (eth *EthereumClassic) CreateRawTx(senderPrivateKey, recipient string, amount int64) (rawTxHex string) {

	client, err := ethclient.Dial("https://ethereumclassic.network")
	if err != nil {
		fmt.Println(err)
		return
	}

	privateKey, err := crypto.HexToECDSA(senderPrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Errorf("error casting public key to ECDSA")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	value := big.NewInt(amount) // in wei (1 eth)
	gasLimit := uint64(21000)   // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	toAddress := common.HexToAddress(recipient)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex = hex.EncodeToString(rawTxBytes)
	return
}

func (eth *EthereumClassic) SendRawTx(rawTx string) {

	client, err := ethclient.Dial("https://ethereumclassic.network")
	if err != nil {
		fmt.Println(err)
		return
	}

	rawTxBytes, err := hex.DecodeString(rawTx)

	tx := new(types.Transaction)

	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
