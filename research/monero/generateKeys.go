package main

import (
	"encoding/hex"
	"fmt"
	"github.com/paxosglobal/moneroutil"
)

func main() {
	prvtKey := moneroutil.RandomScalar()
	fmt.Print("Private:")
	fmt.Println(hex.EncodeToString(prvtKey.Serialize()))
	fmt.Print("Public:")
	fmt.Println(hex.EncodeToString(prvtKey.PubKey().Serialize()))
}
