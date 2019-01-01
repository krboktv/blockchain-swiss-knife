package main

import (
	"encoding/hex"
	"fmt"
	"github.com/krboktv/blockchain-swiss-knife/bitcoinGold"
	"os"

	st "../stellar"
	"encoding/json"
	"github.com/krboktv/blockchain-swiss-knife/bitcoin"
	"github.com/krboktv/blockchain-swiss-knife/dash"
	"github.com/krboktv/blockchain-swiss-knife/ethereum"
	"github.com/krboktv/blockchain-swiss-knife/monero"
	"github.com/krboktv/blockchain-swiss-knife/ripple"
)

func main() {

	//eth()
	//fmt.Print("\n")

	btc()
	fmt.Print("\n")
	//dash_()
	fmt.Print("\n")
	xrp()
	fmt.Println("\n")
	stellar()
	fmt.Println("\n")
	btg()
	fmt.Println("\n")
	xmr()
}

func eth() {

	privateKey, _ := ethereum.GenerateKey()
	publicKey := ethereum.GetPublicKey(privateKey)
	address := ethereum.GetAddress(privateKey)

	balanceTest := ethereum.GetBalance("0x343295B49522CFc38aF517c58eBB78565C42Ed95")

	// create raw tx

	hexTxHash := ethereum.CreateRawTx("privateKey", "recipient", 1000)

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
	fmt.Println("Create Raw TX")
	fmt.Print("txHex:")
	fmt.Println(hexTxHash)
	fmt.Println("---Ethereum---")
}

func btc() {

	privateKey, _ := bitcoin.GenerateKey()
	publicKey := bitcoin.GetPublicKey(privateKey)
	address, _ := bitcoin.GetAddress(privateKey)

	balanceTest := bitcoin.GetBalance("18bXSCSXiTD3DB3XEz851VpB4ZK49rkprT")

	transaction, err := bitcoin.CreateTransaction("5HusYj2b2x4nroApgfvaSfKYZhRbKFH41bVyPooymbC6KfgSXdD", "1KKKK6N21XKo48zWKuQKXdvSsCf95ibHFa", 91234, "81b4c832d70cb56ff957589752eb4125a4cab78a25a8fc52d6a09e5bd4404d48")
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := json.Marshal(transaction)

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
	fmt.Print("Create Raw tx:")
	fmt.Println(string(data))
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
	fmt.Println("Account for balance test: rUjAoB9tXmt5v1DifGnfbDT6WRTX67PXvq")
	fmt.Print("Test Balance: ")
	fmt.Println(balanceTest)
	fmt.Println("---Ripple---")
}

func stellar() {
	x, err := st.GenerateKey()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	balanceTest := st.GetBalance("GAQV4K7OZJMR32NADB3D27DVBIPGDZHLYV3ZOPA57ZS4CCG2QQVUP2UX")

	fmt.Println("---Stellar---")
	fmt.Print("Seed phrase: ")
	fmt.Println(x.Seed)
	fmt.Print("Address: ")
	fmt.Println(x.Address)
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

func xmr() {
	prvtKey := monero.GenerateKey()
	pubKey := monero.GetPublicKey(*prvtKey)
	fmt.Println("---Monero---")
	fmt.Println("Monero private:")
	fmt.Println(hex.EncodeToString(prvtKey.Serialize()))
	fmt.Println("Monero public:")
	fmt.Println(hex.EncodeToString(pubKey.Serialize()))
	fmt.Println("---Monero---")
}