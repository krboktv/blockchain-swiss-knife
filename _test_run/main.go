package main

import (
	st "github.com/krboktv/blockchain-swiss-knife/stellar"
	"encoding/hex"
	"fmt"
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
	eth()
	fmt.Println("\n")
	etc()
	fmt.Println("\n")
	xrp()
	fmt.Println("\n")
	stellar()
	fmt.Println("\n")
	tether()
	fmt.Println("\n")
	_zcash()
}

func eth() {
	swissKnife.Ethereum.GenerateAndSet()

	balanceTest := swissKnife.Ethereum.GetBalance("0x343295B49522CFc38aF517c58eBB78565C42Ed95")

	fmt.Println("---Ethereum---")
	fmt.Print("Private Key: ")
	fmt.Println(swissKnife.Ethereum.PrivateKey)
	fmt.Print("Public Key: ")
	fmt.Println(swissKnife.Ethereum.PublicKey)
	fmt.Print("Address: ")
	fmt.Println(swissKnife.Ethereum.Address)
	fmt.Println("Account for balance test: 0x343295B49522CFc38aF517c58eBB78565C42Ed95")
	fmt.Print("Test Balance: ")
	fmt.Println(balanceTest)
	fmt.Println("---Ethereum---")
}

func etc(){
	swissKnife.EthereumClassic.GenerateAndSet()
	balanceTest := swissKnife.EthereumClassic.GetBalance("0xDf7D7e053933b5cC24372f878c90E62dADAD5d42")

	fmt.Println("---EthereumClassic---")
	fmt.Print("Private Key: ")
	fmt.Println(swissKnife.EthereumClassic.PrivateKey)
	fmt.Print("Public Key: ")
	fmt.Println(swissKnife.EthereumClassic.PublicKey)
	fmt.Print("Address: ")
	fmt.Println(swissKnife.EthereumClassic.Address)
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

	balanceTest := swissKnife.Ripple.GetBalance("rUjAoB9tXmt5v1DifGnfbDT6WRTX67PXvq")

	seed := []byte("sspmdvhjCgmasqzg9a6HW6rvYLEoD")

	seedFromExistingPassphrase, _ := swissKnife.Ripple.GenerateKeyFromPassphrase([]byte("masterpassphrase"))

	childAccount, _ := swissKnife.Ripple.GetChildAccount(seedFromExistingPassphrase)

	publicKey, _ := swissKnife.Ripple.GetPublicKey(seedFromExistingPassphrase)

	pvk, _ := hex.DecodeString("1ACAAEDECE405B2A958212629E16F2EB46B153EEE94CDD350FDEFF52795525B7")

	publicKeyFromPrivateKey := swissKnife.Ripple.GetPublicKeyFromPrivateKey(pvk)

	address, _ := swissKnife.Ripple.GetAddress(seed)

	addressFromPrivateKey, _ := swissKnife.Ripple.GetAddressFromPrivateKey(pvk)

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
	// first check
	seed, _ := swissKnife.Stellar.GenerateKey()

	seedFromExistingPhrase, _ := swissKnife.Stellar.GenerateKeyFromPassphrase([]byte("masterpassphrase"))
	pvk, _ := swissKnife.Stellar.GetPrivateKeyFromSeed(seed)
	pvkHex := hex.EncodeToString(pvk)
	pbk := swissKnife.Stellar.GetPublicKeyFromPrivateKey(pvk)
	pubHex := hex.EncodeToString(pbk)
	address, _ := swissKnife.Stellar.GetAddress(seed)
	addressFromPvk, _ := swissKnife.Stellar.GetAddressFromPrivateKey(pvk)


	// second
	swissKnife.Stellar.GenerateAndSet()
	balanceTest := st.GetBalance("GAQV4K7OZJMR32NADB3D27DVBIPGDZHLYV3ZOPA57ZS4CCG2QQVUP2UX")


	fmt.Println("---Stellar---")
	fmt.Println("First: ")
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
	fmt.Println("\n\n")

	fmt.Println("Second: ")
	fmt.Print("Random seed: ")
	fmt.Println(swissKnife.Stellar.Seed)
	fmt.Print("Private key from seed: ")
	fmt.Println(swissKnife.Stellar.PrivateKey)
	fmt.Print("Public key from private key: ")
	fmt.Println(swissKnife.Stellar.PublicKey)
	fmt.Print("Address from private key: ")
	fmt.Println(swissKnife.Stellar.Address)
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
	swissKnife.Tether.GenerateAndSet()
	balance := swissKnife.Tether.GetBalance("3NrEXrB9qAxXYfRt6jKtBD8QzoU2qtNWDR")

	fmt.Println("---Tether---")
	fmt.Print("Private Key: ")
	fmt.Println(swissKnife.Tether.PrivateKey)
	fmt.Print("Public Key: ")
	fmt.Println(swissKnife.Tether.PublicKey)
	fmt.Print("Address: ")
	fmt.Println(swissKnife.Tether.Address)
	fmt.Println("Account for balance test: 3NrEXrB9qAxXYfRt6jKtBD8QzoU2qtNWDR")
	fmt.Print("Balance:")
	fmt.Println(balance)
	fmt.Println("---Tether---")
}

func _zcash() {
	swissKnife.ZCash.GenerateAndSet()

	fmt.Println("---ZCash---")
	fmt.Print("Private Key: ")
	fmt.Println(swissKnife.ZCash.PrivateKey)
	fmt.Print("Public Key: ")
	fmt.Println(swissKnife.ZCash.PublicKey)
	fmt.Print("Address: ")
	fmt.Println(swissKnife.ZCash.Address)
	fmt.Println("---ZCash---")
}
