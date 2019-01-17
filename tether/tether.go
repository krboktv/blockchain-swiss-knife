package tether

import (
	"encoding/hex"
	"fmt"
	"github.com/imroc/req"
	"github.com/krboktv/blockchain-swiss-knife/utils"
)

type Tether struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

var (
	MainnetTether = []byte{0x00}
)

func (tether *Tether) GenerateKey() ([]byte, error) {
	return utils.GenerateKeySecp256k1()
}

func (tether *Tether) GetPublicKey(key []byte) []byte {
	return utils.GetPublicKeySecp256k1(key)
}

func (tether *Tether) GetAddress(key []byte) ([]byte, error) {
	pbk := tether.GetPublicKey(key)
	step1 := utils.SHA256(pbk)
	step2 := utils.RIPEMD160(step1)
	step3 := append(MainnetTether, step2...)
	step4 := utils.DoubleSHA256(step3)
	step5 := append(step3, step4[:4]...)
	return utils.EncodeToBase58(utils.EncodeBitcoin, step5)
}

func (tether *Tether) GenerateAndSet() {
	privateKey, err := tether.GenerateKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := tether.GetPublicKey(privateKey)
	address, err := tether.GetAddress(privateKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	tether.PrivateKey = hex.EncodeToString(privateKey)
	tether.PublicKey = hex.EncodeToString(publicKey)
	tether.Address = string(address)

}

func (tether *Tether) GetBalance(address string) (balanceStr string) {

	type TetherBalance struct {
		Balance []struct {
			Divisible    bool   `json:"divisible"`
			Frozen       string `json:"frozen"`
			Id           string `json:"id"`
			Pendingneg   string `json:"pendingneg"`
			Pendingpos   string `json:"pendingpos"`
			Propertyinfo struct {
				Amount          string   `json:"amount"`
				Block           int      `json:"block"`
				Blockhash       string   `json:"blockhash"`
				Blocktime       int      `json:"blocktime"`
				Catecory        string   `jpon:"catecory"`
				Confirmations   int      `json:"confirmations"`
				Creationtxid    string   `json:"creationtxid"`
				Data            string   `json:"data"`
				Divisible       bool     `json:"divisible"`
				Ecosystem       string   `json:"ecosystem"`
				Fee             string   `json:"fee"`
				Fixedissuance   bool     `json:"fixedissuance"`
				Flags           struct{} `json:"flags"`
				Freezingenabled bool     `json:"freezingenabled"`
				Ismine          bool     `json:"ismine"`
				Issuer          string   `json:"issuer"`
				Managedissuance bool     `json:"managedissuance"`
				Name            string   `json:"name"`
				Positioninblock int      `json:"positioninblock"`
				Propertyid      int      `json:"propertyid"`
				Propertyname    string   `json:"propertyname"`
				Propertytype    string   `json:"propertytype"`
				Rdata           int      `json:"rdata"`
				Registered      bool     `json:"registered"`
				Sendingaddress  string   `json:"sendingaddress"`
				Subcategory     string   `json:"subcategory"`
				Totaltokens     string   `json:"totaltokens"`
				Txid            string   `json:"txid"`
				Type            string   `json:"type"`
				Type_int        int      `json:"type_int"`
				Url             string   `json:"url"`
				Valid           bool     `json:"valid"`
				Version         int      `json:"version"`
			}
			Reserved string `json:"reserved"`
			Symbol   string `json:"symbol"`
			Value    string `json:"value"`
		}
	}

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	param := req.Param{
		"addr": address,
	}

	reqForBalance, err := req.Post("https://api.omniexplorer.info/v1/address/addr/", header, param)
	if err != nil {
		fmt.Println(err)
		return
	}

	var tb TetherBalance

	reqForBalance.ToJSON(&tb)

	for _, j := range tb.Balance {
		if j.Propertyinfo.Name == "TetherUS" {
			balanceStr = j.Value
		}
	}

	return
}
