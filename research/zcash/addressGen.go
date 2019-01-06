package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/FiloSottile/zcash-mini/zcash"
	"github.com/blackkeyboard/zgenerate/zcashcrypto"
)

func main() {
	// zcash-mini
	fmt.Println("ZCash-mini")
	key := zcash.GenerateKey()
	fmt.Println(hex.EncodeToString(key))

	address, _ := zcash.KeyToAddress(key)
	fmt.Println(hex.EncodeToString(address))

	viewKey, _ := zcash.KeyToViewingKey(key)
	fmt.Println(hex.EncodeToString(viewKey))
	fmt.Println()
	fmt.Println()

	// zgenerate
	fmt.Println("ZGenerate")
	wallet, err := zcashcrypto.CreateWallet(false, 1)
	if err != nil {
		log.Panicln(err.Error())
	}
	log.Println("Wallet generated!")
	log.Printf("Passphrase: %s\n", wallet.Passphrase)
	log.Printf("Address\t\t\t\tPrivate key")

	for i := 0; i <= len(wallet.Addresses)-1; i++ {
		log.Printf("%s\t%s\n", wallet.Addresses[i].Value, wallet.Addresses[i].PrivateKey)
	}
}
