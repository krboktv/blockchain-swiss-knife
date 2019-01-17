package main

import (
	st "github.com/krboktv/blockchain-swiss-knife/stellar"
	"github.com/krboktv/blockchain-swiss-knife/zcash"
	"encoding/hex"
	"fmt"
	"github.com/krboktv/blockchain-swiss-knife/ethereum"
	"github.com/krboktv/blockchain-swiss-knife/ethereumClassic"
	"github.com/krboktv/blockchain-swiss-knife/ripple"
	t "github.com/krboktv/blockchain-swiss-knife/tether"
	. "github.com/krboktv/blockchain-swiss-knife/Knife"
)

var swissKnife Knife

func main() {
	fmt.Println("To the moon!")
}

func init(){
	btc()
	fmt.Println("\n")
	btg()
	fmt.Println("\n")
	dash_()
	fmt.Println("\n")
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
	swissKnife.Bitcoin.GenerateAndSet()
	balanceTest := swissKnife.Bitcoin.GetBalance("18bXSCSXiTD3DB3XEz851VpB4ZK49rkprT")

	fmt.Println("---Bitcoin---")
	fmt.Print("Private Key: ")
	fmt.Println(swissKnife.Bitcoin.PrivateKey)
	fmt.Print("Public Key: ")
	fmt.Println(swissKnife.Bitcoin.PublicKey)
	fmt.Print("Address: ")
	fmt.Println(swissKnife.Bitcoin.Address)
	fmt.Println("Account for balance test: 18bXSCSXiTD3DB3XEz851VpB4ZK49rkprT")
	fmt.Print("Test Balance: ")
	fmt.Println(balanceTest)
	fmt.Println("---Bitcoin---")
}

func dash_() {
	swissKnife.Dash.GenerateAndSet()
	balanceTest := swissKnife.Dash.GetBalance("XkNPrBSJtrHZUvUqb3JF4g5rMB3uzaJfEL")

	fmt.Println("---Dash---")
	fmt.Print("Private Key: ")
	fmt.Println(swissKnife.Dash.PrivateKey)
	fmt.Print("Public Key: ")
	fmt.Println(swissKnife.Dash.PublicKey)
	fmt.Print("Dash Address: ")
	fmt.Println(swissKnife.Dash.Address)
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
	seed, _ := st.GenerateKey()
	seedFromExistingPhrase, _ := st.GenerateKeyFromPassphrase([]byte("masterpassphrase"))
	pvk, _ := st.GetPrivateKeyFromSeed(seed)
	pvkHex := hex.EncodeToString(pvk)
	pbk := st.GetPublicKeyFromPrivateKey(pvk)
	pubHex := hex.EncodeToString(pbk)
	address, _ := st.GetAddress(seed)
	addressFromPvk, _ := st.GetAddressFromPrivateKey(pvk)

	balanceTest := st.GetBalance("GAQV4K7OZJMR32NADB3D27DVBIPGDZHLYV3ZOPA57ZS4CCG2QQVUP2UX")

	fmt.Println("---Stellar---")
	fmt.Print("Random seed: ")
	fmt.Println(string(seed))
	fmt.Print("Seed from existing passphrase: ")
	fmt.Println(string(seedFromExistingPhrase))
	fmt.Print("Private key from seed: ")
	fmt.Println(pvkHex)
	fmt.Print("Public key from private key: ")
	fmt.Println(pubHex)
	fmt.Print("Address from seed: ")
	fmt.Println(string(address))
	fmt.Print("Address from private key: ")
	fmt.Println(string(addressFromPvk))
	fmt.Println("Account for balance test: GAQV4K7OZJMR32NADB3D27DVBIPGDZHLYV3ZOPA57ZS4CCG2QQVUP2UX")
	fmt.Print("Test Balance: ")
	fmt.Println(balanceTest)
	fmt.Println("---Stellar---")
}

func btg() {
	swissKnife.BitcoinGold.GenerateAndSet()

	balanceTest := swissKnife.BitcoinGold.GetBalance("GJjz2Du9BoJQ3CPcoyVTHUJZSj62i1693U")

	fmt.Println("---BitcoinGold---")
	fmt.Print("Private Key: ")
	fmt.Println(swissKnife.BitcoinGold.PrivateKey)
	fmt.Print("Public Key: ")
	fmt.Println(swissKnife.BitcoinGold.PublicKey)
	fmt.Print("Address: ")
	fmt.Println(swissKnife.BitcoinGold.Address)
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
