package main

import (
	"encoding/hex"
	"fmt"
	"github.com/krboktv/blockchain-swiss-knife/bitcoin"
	"github.com/krboktv/blockchain-swiss-knife/dash"
	"github.com/krboktv/blockchain-swiss-knife/ethereum"
	"github.com/krboktv/blockchain-swiss-knife/ripple"
)

func main() {

	eth()
	fmt.Print("\n")
	btc()
	fmt.Print("\n")
	dash_()
	fmt.Print("\n")
	xrp()
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
	address, _ := bitcoin.GetAddress(privateKey)

	fmt.Print("Bitcoin Secret Key: \n")
	fmt.Print(hex.EncodeToString(privateKey))
	fmt.Print("\n")
	fmt.Print("Bitcoin Public Key: \n")
	fmt.Print(hex.EncodeToString(publicKey))
	fmt.Print("\n")
	fmt.Print("Bitcoin Address: \n")
	fmt.Print(string(address))
	fmt.Print("\n")
}

func dash_() {
	privateKey, _ := dash.GenerateKey()
	publicKey := dash.GetPublicKey(privateKey)
	address, _ := dash.GetAddress(privateKey)

	fmt.Print("Dash Secret Key: \n")
	fmt.Print(hex.EncodeToString(privateKey))
	fmt.Print("\n")
	fmt.Print("Dash Public Key: \n")
	fmt.Print(hex.EncodeToString(publicKey))
	fmt.Print("\n")
	fmt.Print("Dash Address: \n")
	fmt.Print(string(address))
	fmt.Print("\n")
}

func xrp() {
	privateKey, _ := ripple.GenerateKey()
	publicKey := ripple.GetPublicKey(privateKey)
	address, _ := ripple.GetAddress(privateKey)

	fmt.Print("Ripple Secret Key: \n")
	fmt.Print(hex.EncodeToString(privateKey))
	fmt.Print("\n")
	fmt.Print("Ripple Public Key: \n")
	fmt.Print(hex.EncodeToString(publicKey))
	fmt.Print("\n")
	fmt.Print("Ripple Address: \n")
	fmt.Print(string(address))
	fmt.Print("\n")
}
