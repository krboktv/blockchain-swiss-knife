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

	//eth()
	//fmt.Print("\n")
	//btc()
	//fmt.Print("\n")
	//dash_()
	//fmt.Print("\n")
	//xrp()
	//fmt.Println("\n")
	stellar()
}

func eth() {
	privateKey, _ := ethereum.GenerateKey()
	publicKey := ethereum.GetPublicKey(privateKey)
	address := ethereum.GetAddress(privateKey)

	fmt.Print("Ethereum Private Key: \n")
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

	fmt.Print("Bitcoin Private Key: \n")
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

	fmt.Print("Dash Private Key: \n")
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
	pvk := ripple.TestGenerateSeed()
	address, _ := ripple.GetAddress(pvk)
	fmt.Print("Ripple Private Key: \n")
	fmt.Print(hex.EncodeToString(pvk))
	fmt.Print("\n")
	fmt.Print("Ripple Public Key: \n")
	fmt.Print(hex.EncodeToString(publicKey))
	fmt.Print("\n")
	fmt.Print("Ripple Address: \n")
	fmt.Print(string(address))
	fmt.Print("\n")
}

func stellar(){
	fmt.Println("Stellar:")

	x,err := st.GeneratePrivateKey()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Public key:")
	fmt.Println(x.Address)
	fmt.Println("Private key:")
	fmt.Println(x.Seed)



}
