package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/krboktv/blockchain-swiss-knife/bitcoin"
	"github.com/krboktv/blockchain-swiss-knife/dash"
	"github.com/krboktv/blockchain-swiss-knife/ethereum"
)

func main() {

	eth()
	fmt.Print("\n")
	btc()
	fmt.Print("\n")
	dash_()
}

func eth() {
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

func btc() {
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

func dash_() {
	privateKey, _ := dash.GenerateKey()
	publicKey := dash.GetPublicKey(privateKey)
	address := dash.GetAddress(privateKey)

	fmt.Print("Dash Secret Key: \n")
	fmt.Print(hex.EncodeToString(privateKey))
	fmt.Print("\n")
	fmt.Print("Dash Public Key: \n")
	fmt.Print(hex.EncodeToString(publicKey))
	fmt.Print("\n")
	fmt.Print("Dash Address: \n")
	fmt.Print(base58.Encode(address))
	fmt.Print("\n")
}
