package ethereum

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/krboktv/blockchain-swiss-knife/utils"
	"github.com/onrik/ethrpc"
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

func GetBalanceForMultipleAdresses(addr []string) map[string]float64 {

	r := New()
	var wg sync.WaitGroup

	for i := 0; i < len(addr); i++ {
		wg.Add(1)
		go worker(&wg, addr[i], r)
	}
	wg.Wait()

	return r.balances
}
