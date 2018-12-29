package main

import (
	"encoding/hex"
	"fmt"
	"os"

	st "../stellar"
	"github.com/krboktv/blockchain-swiss-knife/bitcoin"
	"github.com/krboktv/blockchain-swiss-knife/dash"
	"github.com/krboktv/blockchain-swiss-knife/ethereum"
	"github.com/krboktv/blockchain-swiss-knife/ripple"
	"github.com/krboktv/blockchain-swiss-knife/utils/moneroutil"
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
	btg()
	xmr()
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
	seed := []byte("sspmdvhjCgmasqzg9a6HW6rvYLEoD")
	seedFromExistingPassphrase, _ := ripple.GenerateKeyFromPassphrase([]byte("masterpassphrase"))
	childAccount, _ := ripple.GetChildAccount(seedFromExistingPassphrase)
	publicKey, _ := ripple.GetPublicKey(seedFromExistingPassphrase)
	pvk, _ := hex.DecodeString("1ACAAEDECE405B2A958212629E16F2EB46B153EEE94CDD350FDEFF52795525B7")
	publicKeyFromPrivateKey := ripple.GetPublicKeyFromPrivateKey(pvk)
	address, _ := ripple.GetAddress(seed)
	addressFromPrivateKey, _ := ripple.GetAddressFromPrivateKey(pvk)

	fmt.Println("---Ripple---")
	//fmt.Print("Seed: ")
	//fmt.Println(string(seed))
	fmt.Print("Seed from existing passphrase: ")
	fmt.Println(string(seedFromExistingPassphrase))
	fmt.Print("PrivateKey From Seed: ")
	fmt.Println(hex.EncodeToString(childAccount.GetPrivateKey()))
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

func stellar() {
	fmt.Println("---Stellar---")

	x, err := st.GenerateKey()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print("Seed phrase: ")
	fmt.Println(x.Seed)
	fmt.Print("Address: ")
	fmt.Println(x.Address)
	fmt.Println("---Stellar---")
}

func btg() {
	//privateKey, _ := bitcoinGold.GenerateKey()
	//publicKey := bitcoinGold.GetPublicKey(privateKey)
	//address, _ := bitcoinGold.GetAddress(privateKey)

	fmt.Println("---BitcoinGold---")
	fmt.Print("Private Key: ")
	fmt.Println(hex.EncodeToString(privateKey))
	fmt.Print("Public Key: ")
	fmt.Println(hex.EncodeToString(publicKey))
	fmt.Print("Address: ")
	fmt.Println(string(address))
	fmt.Println("---BitcoinGold---")
}

func xmr() {
	pr, pub := moneroutil.NewKeyPair()
	fmt.Println("Monero private:")
	fmt.Println(pr)
	fmt.Println("Monero public:")
	fmt.Println(pub)
}
