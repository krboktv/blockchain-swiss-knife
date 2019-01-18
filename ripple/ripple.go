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

type Ripple struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

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
	return utils.GetPublicKeySecp256k1(ra.GetPrivateKey())
}

func (ra *RootAccount) GetPrivateKey() []byte {
	return ra.privateKey
}

// TODO: Add custom account index
func (xrp *Ripple) GetChildAccount(seed []byte) (*ChildAccount, error) {
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
	return utils.GetPublicKeySecp256k1(ca.GetPrivateKey())
}

func (xrp *Ripple) GenerateKeyFromPassphrase(passphrase []byte) ([]byte, error) {
	return generateSeedFromPassphrase(passphrase)
}

func (xrp *Ripple) GenerateKey() ([]byte, error) {
	return xrp.generateRandomSeed()
}

func (xrp *Ripple) GetPublicKey(seed []byte) ([]byte, error) {
	return xrp.getPublicKeyFromSeed(seed)
}

func (xrp *Ripple) getPublicKeyFromSeed(key []byte) ([]byte, error) {
	childAccount, err := xrp.GetChildAccount(key)
	return utils.GetPublicKeySecp256k1(childAccount.GetPrivateKey()), err
}

func (xrp *Ripple) GetPublicKeyFromPrivateKey(pvk []byte) []byte {
	return utils.GetPublicKeySecp256k1(pvk)
}

func (xrp *Ripple) GetAddress(seed []byte) ([]byte, error) {
	return xrp.getAddressFromSeed(seed)
}

func (xrp *Ripple) getAddressFromSeed(seed []byte) ([]byte, error) {
	childAccount, err := xrp.GetChildAccount(seed)
	if err != nil {
		return nil, err
	}
	return xrp.GetAddressFromPrivateKey(childAccount.GetPrivateKey())
}

func (xrp *Ripple) GetAddressFromPrivateKey(key []byte) ([]byte, error) {
	networkByte := []byte{0x00}
	pbk := xrp.GetPublicKeyFromPrivateKey(key)
	step1 := utils.SHA256(pbk)
	step2 := utils.RIPEMD160(step1)
	step3 := append(networkByte, step2...)
	step4 := utils.SHA256(step3)
	step5 := utils.SHA256(step4)
	step6 := step5[:4]
	step7 := append(step3, step6...)
	return utils.EncodeToBase58(utils.EncodeRipple, step7)
}

func (xrp *Ripple) generateRandomPassphrase() ([]byte, error) {
	phrase := make([]byte, 16)
	rnd, err := rand.Read(phrase)
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint32(phrase, uint32(rnd))
	return phrase, nil
}

func (xrp *Ripple) generateRandomSeed() ([]byte, error) {
	passphrase, err := xrp.generateRandomPassphrase()
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

//func(xrp *Ripple) SetPassphrase(passphrase string){
//	xrp.Passphrase = []byte(passphrase)
//}
//
//func(xrp *Ripple) GetPassphrase() string {
//	return string(xrp.Passphrase)
//}

// Generate acc
// todo: add GenerateAndSet

//func (xrp *Ripple) GenerateAndSet(passphrase string) {
//	seedFromExistingPassphrase, err := xrp.GenerateKeyFromPassphrase([]byte(passphrase))
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//
//	publicKey,err := xrp.GetPublicKey(seedFromExistingPassphrase)
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//
//}

func (xrp *Ripple) GetBalance(address string) (balanceFloat float64) {

	type RippleBalance struct {
		Balances []struct {
			Currency string `json:"currency"`
			Value    string `json:"value"`
		}
	}

	balance, err := req.Get("https://data.ripple.com/v2/accounts/" + address + "/balances?currency=XRP")
	if err != nil {
		fmt.Println(err)
		return
	}

	var b RippleBalance
	balance.ToJSON(&b)

	balanceFloat, err = strconv.ParseFloat(b.Balances[0].Value, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
