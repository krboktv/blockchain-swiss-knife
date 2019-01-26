package test

import (
	"testing"
	. "github.com/krboktv/blockchain-swiss-knife/Knife"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/magiconair/properties/assert"
	"regexp"
)

var swissKnife Knife

func TestETH(t *testing.T) {

	// is valid private key
	rePrivate := regexp.MustCompile("^0x[0-9a-fA-F]{64}$")

	privateKey, err := swissKnife.Ethereum.GenerateKey()
	if err != nil{
		fmt.Println(err)
		return
	}

	assert.Equal(t, rePrivate.MatchString(hexutil.Encode(privateKey)), true)

	// is valid public key
	rePublic := regexp.MustCompile("^[0-9a-fA-F]{128}$")

	public := swissKnife.Ethereum.GetPublicKey(privateKey)

	assert.Equal(t, rePublic.MatchString(hexutil.Encode(public)[4:]),true)


	// is valid address
	reAddress := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	address := hexutil.Encode(swissKnife.Ethereum.GetAddress(privateKey))

	assert.Equal(t, reAddress.MatchString(address), true)

}

