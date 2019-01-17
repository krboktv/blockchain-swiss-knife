package ethereum

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"sync"

	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/krboktv/blockchain-swiss-knife/utils"
	"github.com/onrik/ethrpc"
)

type Ethereum struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

func (eth *Ethereum) GenerateKey() ([]byte, error) {
	return utils.GenerateKeySecp256k1()
}

func (eth *Ethereum) GetPublicKey(key []byte) []byte {
	return utils.GetPublicKeyUncompressedSecp256k1(key)
}

func (eth *Ethereum) GetAddress(key []byte) []byte {
	pbk := eth.GetPublicKey(key)
	return utils.Keccak256(pbk[1:])[12:32]
}

func (eth *Ethereum) GenerateAndSet() {
	privateKey, err := eth.GenerateKey()
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := eth.GetPublicKey(privateKey)
	address := eth.GetAddress(privateKey)

	eth.PrivateKey = hex.EncodeToString(privateKey)
	eth.PublicKey = hex.EncodeToString(publicKey)
	eth.Address = hex.EncodeToString(address)
}

func (eth *Ethereum) GetBalance(address string) (balanceFloat float64) {

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

// Balance by addresses list

type Data struct {
	sync.Mutex
	balances map[string]float64
}

func New() *Data {
	return &Data{
		balances: make(map[string]float64),
	}
}

func (ds *Data) set(key string, value float64) {
	ds.balances[key] = value
}

func (ds *Data) Set(key string, value float64) {
	ds.Lock()
	defer ds.Unlock()
	ds.set(key, value)
}

func worker(wg *sync.WaitGroup, addr string, r *Data) {
	defer wg.Done()
	ethClient := ethrpc.New("https://mainnet.infura.io")
	balance, err := ethClient.EthGetBalance(addr, "latest")
	if err != nil {
		fmt.Println(err)
	}
	floatBalance, _ := strconv.ParseFloat(balance.String(), 64)
	ethBalance := floatBalance / math.Pow(10, 18)
	r.Set(addr, ethBalance)
}

func (eth *Ethereum) GetBalanceForMultipleAdresses(addr []string) map[string]float64 {

	r := New()
	var wg sync.WaitGroup

	for i := 0; i < len(addr); i++ {
		wg.Add(1)
		go worker(&wg, addr[i], r)
	}
	wg.Wait()

	return r.balances
}
