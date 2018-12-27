package ripple

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/krboktv/blockchain-swiss-knife/utils"
)

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

func GetPrivateKeyFromSeed(seed []byte) ([]byte, error) {
	fourZeroBytes := []byte{0x00, 0x00, 0x00, 0x00}
	seedBytes, err := utils.DecodeFromBase58(utils.EncodeRipple, seed)
	if err != nil {
		return nil, err
	}
	update := append(seedBytes[1:17], fourZeroBytes...)
	hash := utils.SHA512(update)
	return hash[:32], nil
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
	pvk, err := GetPrivateKeyFromSeed(key)
	return utils.GetPublicKey(pvk), err
}

func GetPublicKeyFromPrivateKey(pvk []byte) []byte {
	return utils.GetPublicKey(pvk)
}

func GetAddress(seed []byte) ([]byte, error) {
	return getAddressFromSeed(seed)
}

func getAddressFromSeed(seed []byte) ([]byte, error) {
	pvk, err := GetPrivateKeyFromSeed(seed)
	if err != nil {
		return nil, err
	}
	return GetAddressFromPrivateKey(pvk)
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