package ripple

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/imroc/req"
	"github.com/krboktv/blockchain-swiss-knife/utils"
	"strconv"
)

type RootAccount struct {
	privateKey []byte
}

type ChildAccount struct {
	index      []byte
	privateKey []byte
}

func GetFamilyGenerator(seed []byte) (*RootAccount, error) {
	curveOrderBytes, _ := hex.DecodeString(utils.CurveOrder)
	curveOrderUint32 := binary.BigEndian.Uint32(curveOrderBytes)
	counter := uint32(0)

	seedBytes, err := utils.DecodeFromBase58(utils.EncodeRipple, seed)
	if err != nil {
		return nil, err
	}

	var pvk []byte
	for check := true; check; check = binary.BigEndian.Uint32(pvk) > uint32(curveOrderUint32) {
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, counter)
		update := append(seedBytes[1:17], buf.Bytes()...)
		pvk = utils.SHA512(update)[:32]
		counter++
	}

	return &RootAccount{
		pvk,
	}, nil
}

func (ra *RootAccount) GetPublicKey() []byte {
	return utils.GetPublicKey(ra.GetPrivateKey())
}

func (ra *RootAccount) GetPrivateKey() []byte {
	return ra.privateKey
}

// TODO: Add custom account index
func GetChildAccount(seed []byte) (*ChildAccount, error) {
	accountIndex := []byte{0x00, 0x00, 0x00, 0x00}

	curveOrderBytes, _ := hex.DecodeString(utils.CurveOrder)
	curveOrderUint32 := binary.BigEndian.Uint32(curveOrderBytes)
	counter := uint32(0)

	rootAccount, err := GetFamilyGenerator(seed)
	if err != nil {
		return nil, err
	}

	var update1 []byte
	for check := true; check; check = binary.BigEndian.Uint32(update1) > uint32(curveOrderUint32) {
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, counter)
		update := append(append(rootAccount.GetPublicKey(), accountIndex...), buf.Bytes()...)
		update1 = utils.AddBytes(rootAccount.GetPrivateKey(), utils.SHA512(update)[:32]).Bytes()
		counter++
	}

	return &ChildAccount{
		accountIndex,
		utils.ModBytes(update1, curveOrderBytes).Bytes(),
	}, nil
}

func (ca *ChildAccount) GetAccountIndex() []byte {
	return ca.index
}

func (ca *ChildAccount) GetPrivateKey() []byte {
	return ca.privateKey
}

func (ca *ChildAccount) GetPublicKey() []byte {
	return utils.GetPublicKey(ca.GetPrivateKey())
}

func GenerateKeyFromPassphrase(passphrase []byte) ([]byte, error) {
	return generateSeedFromPassphrase(passphrase)
}

func GenerateKey() ([]byte, error) {
	return generateRandomSeed()
}

func GetPublicKey(seed []byte) ([]byte, error) {
	return getPublicKeyFromSeed(seed)
}

func getPublicKeyFromSeed(key []byte) ([]byte, error) {
	childAccount, err := GetChildAccount(key)
	return utils.GetPublicKey(childAccount.GetPrivateKey()), err
}

func GetPublicKeyFromPrivateKey(pvk []byte) []byte {
	return utils.GetPublicKey(pvk)
}

func GetAddress(seed []byte) ([]byte, error) {
	return getAddressFromSeed(seed)
}

func getAddressFromSeed(seed []byte) ([]byte, error) {
	childAccount, err := GetChildAccount(seed)
	if err != nil {
		return nil, err
	}
	return GetAddressFromPrivateKey(childAccount.GetPrivateKey())
}

func GetAddressFromPrivateKey(key []byte) ([]byte, error) {
	networkByte := []byte{0x00}
	pbk := GetPublicKeyFromPrivateKey(key)
	step1 := utils.SHA256(pbk)
	step2 := utils.RIPEMD160(step1)
	step3 := append(networkByte, step2...)
	step4 := utils.SHA256(step3)
	step5 := utils.SHA256(step4)
	step6 := step5[:4]
	step7 := append(step3, step6...)
	return utils.EncodeToBase58(utils.EncodeRipple, step7)
}

func generateRandomPassphrase() ([]byte, error) {
	phrase := make([]byte, 16)
	rnd, err := rand.Read(phrase)
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint32(phrase, uint32(rnd))
	return phrase, nil
}

func generateRandomSeed() ([]byte, error) {
	passphrase, err := generateRandomPassphrase()
	if err != nil {
		return nil, err
	}
	seed := utils.SHA512(passphrase)[:16]
	return encodeSeedToBase58Check(seed)
}

func generateSeedFromPassphrase(passphrase []byte) ([]byte, error) {
	seed := utils.SHA512(passphrase)[:16]
	return encodeSeedToBase58Check(seed)
}

func encodeSeedToBase58Check(seed []byte) ([]byte, error) {
	networkByte := []byte{0x21}
	step1 := append(networkByte, seed...)
	step2 := utils.DoubleSHA256(step1)
	step3 := append(step1, step2[:4]...)
	step4, err := utils.EncodeToBase58(utils.EncodeRipple, step3)
	return step4, err
}

func GetBalance(address string) (balanceFloat float64) {

	type RippleBalance struct {
		Balance_changes []struct {
			Amount_change string `json:"amount_change"`
			Final_balance string `json:"final_balance"`
			Node_index    int    `json:"node_index"`
			Tx_index      int    `json:"tx_index"`
			Change_type   string `json:"change_type"`
			Currency      string `json:"currency"`
			Executed_time string `json:"executed_time"`
			Ledger_index  int    `json:"ledger_index"`
			Tx_hash       string `json:"tx_hash"`
		}
	}

	balance, err := req.Get("https://data.ripple.com/v2/accounts/" + address + "/balance_changes?descending=true&limit=1")
	if err != nil {
		fmt.Println(err)
	}

	var b RippleBalance
	balance.ToJSON(&b)

	balanceFloat, err = strconv.ParseFloat(b.Balance_changes[0].Final_balance, 64)
	if err != nil {
		fmt.Println(err)
	}
	return

}
