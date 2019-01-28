package test

import (
	"testing"
	. "github.com/krboktv/blockchain-swiss-knife/Knife"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/magiconair/properties/assert"
	"regexp"
	"log"
)

var swissKnife Knife


// ETH and ETC
func TestETH(t *testing.T) {

	// is valid private key
	rePrivate := regexp.MustCompile("^0x[0-9a-fA-F]{64}$")

	privateKey, err := swissKnife.Ethereum.GenerateKey()
	if err != nil{
		log.Fatal(err)
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

// BTC, Tether, ZCash, Dash, BTG
func TestBTC(t *testing.T){

	// is valid private key
	rePrivate := regexp.MustCompile("^0x[0-9a-fA-F]{64}$")

	privateKey, err := swissKnife.Bitcoin.GenerateKey()
	if err != nil{
		log.Fatal(err)
	}

	assert.Equal(t, rePrivate.MatchString(hexutil.Encode(privateKey)) ,true)

	// is valid public key
	rePublic := regexp.MustCompile("^[0-9a-fA-F]{66}$")

	publicKey := swissKnife.Bitcoin.GetPublicKey(privateKey)

	assert.Equal(t, rePublic.MatchString(hexutil.Encode(publicKey)[2:]) ,true)

	// is valid address
	reAddress := regexp.MustCompile("^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$")

	address, err := swissKnife.Bitcoin.GetAddress(privateKey)
	if err != nil{
		log.Fatal(err)
	}

	assert.Equal(t, reAddress.MatchString(string(address)) ,true)
}

func TestStellar(t *testing.T){

	// is valid seed
	reSeed := regexp.MustCompile("^[0-9-A-Z]{56}$")

	seed,err := swissKnife.Stellar.GenerateKey()
	if err != nil{
		log.Fatal(err)
	}

	assert.Equal(t, reSeed.MatchString(string(seed)),true)

	// is valid address
	reAddress := regexp.MustCompile("^[0-9-A-Z]{56}$")

	address,err := swissKnife.Stellar.GetAddress(seed)
	if err != nil{
		log.Fatal(err)
	}

	assert.Equal(t, reAddress.MatchString(string(address)),true)

	// get address func
	publicKey := "GDTCMJ4FJNY2SOEJRXUMAR262L7FQKJP5MQHMZVSDGPUN6U7JEORVPQP"
	secretKey := []byte(`SAK5JNDTZ3HZAXZSQFMYYU5OC3JA7PMOITNEGBJAV635BL7B7R2OQAC5`)

	address,err = swissKnife.Stellar.GetAddress(secretKey)
	if err != nil{
		log.Fatal(err)
	}

	assert.Equal(t, string(address), publicKey)

}


func TestRipple(t *testing.T){

	// is valid key
	reSecret := regexp.MustCompile("^s[a-zA-Z-0-9]{28}$")

	seed, _ := swissKnife.Ripple.GenerateKeyFromPassphrase([]byte("999999999"))

	assert.Equal(t, reSecret.MatchString(string(seed)),true)

	// is valid address
	reAddress := regexp.MustCompile("^r[a-zA-Z-0-9]{24,34}$")

	address, _ := swissKnife.Ripple.GetAddress(seed)

	assert.Equal(t, reAddress.MatchString(string(address)),true)

	// get secret and address

	// passphrase - masterpassphrase
	sec := "snoPBrXtMeMyMHUVTgbuqAfg1SUTb"
	addr  := "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh"

	seedFromExistingPassphrase, _ := swissKnife.Ripple.GenerateKeyFromPassphrase([]byte("masterpassphrase"))

	addressFromSeed, _ := swissKnife.Ripple.GetAddress(seedFromExistingPassphrase)

	assert.Equal(t, string(seedFromExistingPassphrase), sec)
	assert.Equal(t, string(addressFromSeed), addr)
}

