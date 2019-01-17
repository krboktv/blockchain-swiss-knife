package stellar

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/imroc/req"
	"github.com/krboktv/blockchain-swiss-knife/utils"
	"github.com/stellar/go/crc16"
	"strconv"
)

const (
	VersionByteSeed = 18 << 3
	VersionByteAccountID = 6 << 3
)

type VersionByte byte

func GenerateKey() ([]byte, error) {
	return generateRandomSeed()
}

func GetPrivateKeyFromSeed(seed []byte) ([]byte, error) {
	seed, err := MustDecode(seed)
	if err != nil {
		return nil, err
	}
	pvk, _, err := utils.GenerateKeyEd25519FromSeed(seed)
	return pvk, err
}

func GenerateKeyFromPassphrase(passphrase []byte) ([]byte, error) {
	return generateSeedFromPassphrase(passphrase)
}

func GetPublicKeyFromPrivateKey(pvk []byte) ([]byte) {
	return utils.GetPublicKeyEd25519(pvk)
}

func GetAddress(seed []byte) ([]byte, error) {
	return getAddressFromSeed(seed)
}

func GetAddressFromPublicKey(pbk []byte) ([]byte, error) {
	return Encode(VersionByteAccountID, pbk)
}

func GetAddressFromPrivateKey(pvt []byte) ([]byte, error) {
	pbk := GetPublicKeyFromPrivateKey(pvt)
	return Encode(VersionByteAccountID, pbk)
}

func GetPublicKey(seed []byte) ([]byte, error) {
	return getPublicKeyFromSeed(seed)
}

func getPublicKeyFromSeed(seed []byte) ([]byte, error) {
	seed, _ = MustDecode(seed)
	_, pbk, err := utils.GenerateKeyEd25519FromSeed(seed)
	return pbk, err
}

func getAddressFromSeed(seed []byte) ([]byte, error) {
	seed, _ = MustDecode(seed)
	_, pbk, err := utils.GenerateKeyEd25519FromSeed(seed)
	if err != nil {
		return nil, err
	}
	return GetAddressFromPublicKey(pbk)
}

func generateRandomPassphrase() ([]byte, error) {
	phrase := make([]byte, 32)
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
	seed := utils.SHA512(passphrase)[:32]
	encodedSeed, err := Encode(VersionByteSeed, seed)
	return encodedSeed, err
}

func generateSeedFromPassphrase(passphrase []byte) ([]byte, error) {
	seed := utils.SHA512(passphrase)[:32]
	return Encode(VersionByteSeed, seed)
}

func Encode(version VersionByte, src []byte) ([]byte, error) {
	var raw bytes.Buffer

	if err := binary.Write(&raw, binary.LittleEndian, version); err != nil {
		return nil, err
	}

	if _, err := raw.Write(src); err != nil {
		return nil, err
	}

	checksum := utils.CRC16Checksum(raw.Bytes())
	if _, err := raw.Write(checksum); err != nil {
		return nil, err
	}

	result := utils.EncodeToBase32(raw.Bytes())
	return []byte(result), nil
}

func MustDecode(seed []byte) ([]byte, error) {
	if len(seed) != 32 {
		var err error
		seed, err = Decode(string(seed))
		if err != nil {
			return nil, err
		}
	}
	return seed, nil
}

func Decode(src string) ([]byte, error) {
	raw, err := utils.DecodeFromBase32(src)
	if err != nil {
		return nil, err
	}

	vp := raw[:len(raw)-2]
	payload := raw[1:len(raw)-2]

	checksum := raw[len(raw)-2:]

	if err := crc16.Validate(vp, checksum); err != nil {
		return nil, err
	}

	return payload, nil
}

func GetBalance(address string) (balanceFloat float64) {

	type StellarBalance struct {
		Balances []struct {
			Balance             string `json:"balance"`
			Buying_liabilities  string `json:"buying_liabilities"`
			Selling_liabilities string `json:"selling_liabilities"`
			Asset_type          string `json:"asset_type"`
		}
	}

	var b StellarBalance
	balance, err := req.Get("https://horizon.stellar.org/accounts/" + address)
	if err != nil {
		fmt.Println(err)
	}

	balance.ToJSON(&b)

	var stellarBalanceString string

	for _, j := range b.Balances {
		if j.Asset_type == "native" {

			stellarBalanceString = j.Balance
		}
	}

	balanceFloat, err = strconv.ParseFloat(stellarBalanceString, 64)
	if err != nil {
		fmt.Println(err)
	}

	return
}
