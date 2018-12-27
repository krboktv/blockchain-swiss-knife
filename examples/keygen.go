package main

import (
	"encoding/hex"
	"fmt"
	"github.com/krboktv/blockchain-swiss-knife/bitcoin"
	"github.com/krboktv/blockchain-swiss-knife/dash"
	"github.com/krboktv/blockchain-swiss-knife/ethereum"
	"github.com/krboktv/blockchain-swiss-knife/ripple"
	st "../stellar"
	"os"
)

func main() {

	eth()
	fmt.Print("\n")
	btc()
	fmt.Print("\n")
	dash_()
	fmt.Print("\n")
	xrp()
	fmt.Println("\n")
	stellar()
	fmt.Println("\n")
}

func eth() {
	privateKey, _ := ethereum.GenerateKey()
	publicKey := ethereum.GetPublicKey(privateKey)
	address := ethereum.GetAddress(privateKey)

	fmt.Println("---Ethereum---")
	fmt.Print("Private Key: ")
	fmt.Println(hex.EncodeToString(privateKey))
	fmt.Print("Public Key: ")
	fmt.Println(hex.EncodeToString(publicKey))
	fmt.Print("Address: ")
	fmt.Println(hex.EncodeToString(address))
	fmt.Println("---Ethereum---")
}

func btc() {
	privateKey, _ := bitcoin.GenerateKey()
	publicKey := bitcoin.GetPublicKey(privateKey)
	address, _ := bitcoin.GetAddress(privateKey)

	fmt.Println("---Bitcoin---")
	fmt.Print("Private Key: ")
	fmt.Println(hex.EncodeToString(privateKey))
	fmt.Print("Public Key: ")
	fmt.Println(hex.EncodeToString(publicKey))
	fmt.Print("Address: ")
	fmt.Println(string(address))
	fmt.Println("---Bitcoin---")

}

func dash_() {
	privateKey, _ := dash.GenerateKey()
	publicKey := dash.GetPublicKey(privateKey)
	address, _ := dash.GetAddress(privateKey)

	fmt.Println("---Dash---")
	fmt.Print("Private Key: ")
	fmt.Println(hex.EncodeToString(privateKey))
	fmt.Print("Public Key: ")
	fmt.Println(hex.EncodeToString(publicKey))
	fmt.Print("Dash Address: ")
	fmt.Println(string(address))
	fmt.Println("---Dash---")
}

func xrp() {
	seed, _ := ripple.GenerateKey()
	seedFromExistingPassphrase, _ := ripple.GenerateKeyFromPassphrase([]byte("masterpassphrase"))
	privateKey, _ := ripple.GetPrivateKeyFromSeed(seedFromExistingPassphrase)
	publicKey, _ := ripple.GetPublicKey(seed)
	pvk, _ := hex.DecodeString("1ACAAEDECE405B2A958212629E16F2EB46B153EEE94CDD350FDEFF52795525B7")
	publicKeyFromPrivateKey := ripple.GetPublicKeyFromPrivateKey(pvk)
	address, _ := ripple.GetAddress(seed)
	addressFromPrivateKey, _ := ripple.GetAddressFromPrivateKey(pvk)

	fmt.Println("---Ripple---")
	fmt.Print("Seed: ")
	fmt.Println(string(seed))
	fmt.Print("Seed from existing passphrase: ")
	fmt.Println(string(seedFromExistingPassphrase))
	fmt.Print("PrivateKey From Seed: ")
	fmt.Println(hex.EncodeToString(privateKey))
	fmt.Print("Public Key: ")
	fmt.Println(hex.EncodeToString(publicKey))
	fmt.Print("Public Key From Private Key: ")
	fmt.Println(hex.EncodeToString(publicKeyFromPrivateKey))
	fmt.Print("Address From Seed: ")
	fmt.Println(string(address))
	fmt.Print("Address From Private Key: ")
	fmt.Println(string(addressFromPrivateKey))
	fmt.Println("---Ripple---")
}

func stellar(){
	fmt.Println("---Stellar---")

	x,err := st.GenerateKey()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print("Public key: ")
	fmt.Print(x.Address)
	fmt.Println()
	fmt.Print("Private key: ")
	fmt.Println(x.Seed)
	fmt.Println("---Stellar---")




}
