package main

import (
	"github.com/paxosglobal/moneroutil"
	"fmt"
	"encoding/hex"
)

func main(){
	prvtKey := moneroutil.RandomScalar()
	fmt.Print("Private:")
	fmt.Println(hex.EncodeToString(prvtKey.Serialize()))
	fmt.Print("Public:")
	fmt.Println(hex.EncodeToString(prvtKey.PubKey().Serialize()))
}