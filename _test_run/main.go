package main

import (
	"encoding/hex"
	"fmt"

	"github.com/krboktv/blockchain-swiss-knife/bitcoinGold"

	st "../stellar"
	"github.com/krboktv/blockchain-swiss-knife/bitcoin"
	"github.com/krboktv/blockchain-swiss-knife/dash"
	"github.com/krboktv/blockchain-swiss-knife/ethereum"
	"github.com/krboktv/blockchain-swiss-knife/ripple"
	t "github.com/krboktv/blockchain-swiss-knife/tether"
	"github.com/krboktv/blockchain-swiss-knife/ethereumClassic"
	"../zcash"
)

func main() {

	btc()
	fmt.Print("\n")
	dash_()
	fmt.Print("\n")
	xrp()
	fmt.Println("\n")
	stellar()
	fmt.Println("\n")
	btg()
	fmt.Println("\n")
	eth()
	fmt.Println("\n")
	etc()
	fmt.Println("\n")
	tether()
	fmt.Println("\n")
	_zcash()

}

func eth() {
	privateKey, _ := ethereum.GenerateKey()
	publicKey := ethereum.GetPublicKey(privateKey)
	address := ethereum.GetAddress(privateKey)
	balanceTest := ethereum.GetBalance("0x343295B49522CFc38aF517c58eBB78565C42Ed95")

	fmt.Println("---Ethereum---")
	fmt.Print("Private Key: ")
	fmt.Println(hex.EncodeToString(privateKey))
	fmt.Print("Public Key: ")
	fmt.Println(hex.EncodeToString(publicKey))
	fmt.Print("Address: ")
	fmt.Println(hex.EncodeToString(address))
	fmt.Println("Account for balance test: 0x343295B49522CFc38aF517c58eBB78565C42Ed95")
	fmt.Print("Test Balance: ")
	fmt.Println(balanceTest)
	fmt.Println("---Ethereum---")
}

func etc(){
	privateKey, _ := ethereumClassic.GenerateKey()
	publicKey := ethereumClassic.GetPublicKey(privateKey)
	address := ethereumClassic.GetAddress(privateKey)
	balanceTest := ethereumClassic.GetBalance("0xDf7D7e053933b5cC24372f878c90E62dADAD5d42")

	fmt.Println("---EthereumClassic---")
	fmt.Print("Private Key: ")
	fmt.Println(hex.EncodeToString(privateKey))
	fmt.Print("Public Key: ")
	fmt.Println(hex.EncodeToString(publicKey))
	fmt.Print("Address: ")
	fmt.Println(hex.EncodeToString(address))
	fmt.Println("Account for balance test: 0xDf7D7e053933b5cC24372f878c90E62dADAD5d42")
	fmt.Print("Test Balance: ")
	fmt.Println(balanceTest)
	fmt.Println("---EthereumClassic---")
}


func btc() {
	privateKey, _ := bitcoin.GenerateKey()
	publicKey := bitcoin.GetPublicKey(privateKey)
	address, _ := bitcoin.GetAddress(privateKey)
	balanceTest := bitcoin.GetBalance("18bXSCSXiTD3DB3XEz851VpB4ZK49rkprT")

	fmt.Println("---Bitcoin---")
	fmt.Print("Private Key: ")
	fmt.Println(hex.EncodeToString(privateKey))
	fmt.Print("Public Key: ")
	fmt.Println(hex.EncodeToString(publicKey))
	fmt.Print("Address: ")
	fmt.Println(string(address))
	fmt.Println("Account for balance test: 18bXSCSXiTD3DB3XEz851VpB4ZK49rkprT")
	fmt.Print("Test Balance: ")
	fmt.Println(balanceTest)
	fmt.Println("---Bitcoin---")
}

func dash_() {
	privateKey, _ := dash.GenerateKey()
	publicKey := dash.GetPublicKey(privateKey)
	address, _ := dash.GetAddress(privateKey)
	balanceTest := dash.GetBalance("XkNPrBSJtrHZUvUqb3JF4g5rMB3uzaJfEL")

	fmt.Println("---Dash---")
	fmt.Print("Private Key: ")
	fmt.Println(hex.EncodeToString(privateKey))
	fmt.Print("Public Key: ")
	fmt.Println(hex.EncodeToString(publicKey))
	fmt.Print("Dash Address: ")
	fmt.Println(string(address))
	fmt.Println("Account for balance test: XkNPrBSJtrHZUvUqb3JF4g5rMB3uzaJfEL")
	fmt.Print("Test Balance: ")
	fmt.Println(balanceTest)
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
	balanceTest := ripple.GetBalance("rUjAoB9tXmt5v1DifGnfbDT6WRTX67PXvq")

	fmt.Println("---Ripple---")
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
	fmt.Println("Account for balance test: rUjAoB9tXmt5v1DifGnfbDT6WRTX67PXvq")
	fmt.Print("Test Balance: ")
	fmt.Println(balanceTest)
	fmt.Println("---Ripple---")
}

func stellar() {
	balanceTest := st.GetBalance("GAQV4K7OZJMR32NADB3D27DVBIPGDZHLYV3ZOPA57ZS4CCG2QQVUP2UX")

	fmt.Println("---Stellar---")
	fmt.Println("Account for balance test: GAQV4K7OZJMR32NADB3D27DVBIPGDZHLYV3ZOPA57ZS4CCG2QQVUP2UX")
	fmt.Print("Test Balance: ")
	fmt.Println(balanceTest)
	fmt.Println("---Stellar---")
}

func btg() {
	privateKey, _ := bitcoinGold.GenerateKey()
	publicKey := bitcoinGold.GetPublicKey(privateKey)
	address, _ := bitcoinGold.GetAddress(privateKey)
	balanceTest := bitcoinGold.GetBalance("GJjz2Du9BoJQ3CPcoyVTHUJZSj62i1693U")

	fmt.Println("---BitcoinGold---")
	fmt.Print("Private Key: ")
	fmt.Println(hex.EncodeToString(privateKey))
	fmt.Print("Public Key: ")
	fmt.Println(hex.EncodeToString(publicKey))
	fmt.Print("Address: ")
	fmt.Println(string(address))
	fmt.Println("Account for balance test: GJjz2Du9BoJQ3CPcoyVTHUJZSj62i1693U")
	fmt.Print("Test Balance: ")
	fmt.Println(balanceTest)
	fmt.Println("---BitcoinGold---")
}

func tether() {
	privateKey, _ := t.GenerateKey()
	publicKey := t.GetPublicKey(privateKey)
	address, _ := t.GetAddress(privateKey)
	balance := t.GetBalance("3NrEXrB9qAxXYfRt6jKtBD8QzoU2qtNWDR")

	fmt.Println("---Tether---")
	fmt.Print("Private Key: ")
	fmt.Println(hex.EncodeToString(privateKey))
	fmt.Print("Public Key: ")
	fmt.Println(hex.EncodeToString(publicKey))
	fmt.Print("Address: ")
	fmt.Println(string(address))
	fmt.Println("Account for balance test: 3NrEXrB9qAxXYfRt6jKtBD8QzoU2qtNWDR")
	fmt.Print("Balance:")
	fmt.Println(balance)
	fmt.Println("---Tether---")
}

func _zcash() {
	privateKey, _ := zcash.GenerateKey()
	publicKey := zcash.GetPublicKey(privateKey)
	address, _ := zcash.GetAddress(privateKey)

	fmt.Println("---ZCash---")
	fmt.Print("Private Key: ")
	fmt.Println(hex.EncodeToString(privateKey))
	fmt.Print("Public Key: ")
	fmt.Println(hex.EncodeToString(publicKey))
	fmt.Print("Address: ")
	fmt.Println(string(address))
	fmt.Println("---ZCash---")
}