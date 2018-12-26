package main

import (
	"encoding/hex"
	"fmt"
	"github.com/krboktv/blockchain-swiss-knife/ethereum"
	"github.com/krboktv/blockchain-swiss-knife/bitcoin"
	"github.com/btcsuite/btcutil/base58"
)

func main()  {

	eth()
	fmt.Print("\n")
	btc()

}

func eth()  {
	privateKey, _ := ethereum.GenerateKey()
	publicKey := ethereum.GetPublicKey(privateKey)
	address := ethereum.GetAddress(privateKey)

	fmt.Print("Ethereum Secret Key: \n")
	fmt.Print(hex.EncodeToString(privateKey))
	fmt.Print("\n")
	fmt.Print("Ethereum Public Key: \n")
	fmt.Print(hex.EncodeToString(publicKey))
	fmt.Print("\n")
	fmt.Print("Ethereum Address: \n")
	fmt.Print(hex.EncodeToString(address))
	fmt.Print("\n")
}

func btc()  {
	privateKey, _ := bitcoin.GenerateKey()
	publicKey := bitcoin.GetPublicKey(privateKey)
	address := bitcoin.GetAddress(privateKey)

	fmt.Print("Bitcoin Secret Key: \n")
	fmt.Print(hex.EncodeToString(privateKey))
	fmt.Print("\n")
	fmt.Print("Bitcoin Public Key: \n")
	fmt.Print(hex.EncodeToString(publicKey))
	fmt.Print("\n")
	fmt.Print("Bitcoin Address: \n")
	fmt.Print(base58.Encode(address))
	fmt.Print("\n")
}
