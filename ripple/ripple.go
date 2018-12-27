package ripple

import (
	"crypto/rand"
	"github.com/krboktv/blockchain-swiss-knife/utils"
)

func generateRandomPassphrase() []byte {
	phrase := make([]byte, 16)
	rand.Read(phrase)
	return phrase
}

func generateRandomSeed() ([]byte, error) {
	passphrase := generateRandomPassphrase()
	seed := utils.SHA512(passphrase)[:16]
	return encodeSeedToBase58Check(seed)
}

func generateSeedFromPassphrase(passphrase string) ([]byte, error) {
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

func GenerateKeyFromPassphrase(passphrase string) ([]byte, error) {
	return generateSeedFromPassphrase(passphrase)
}

func GenerateKey() ([]byte, error) {
	return generateRandomSeed()
}

func GetPublicKey(key []byte) []byte {
	return utils.GetPublicKey(key)
}

func GetAddress(key []byte) ([]byte, error) {
	networkByte := []byte{0x00}
	pbk := GetPublicKey(key)
	step1 := utils.SHA256(pbk)
	step2 := utils.RIPEMD160(step1)
	step3 := append(networkByte, step2...)
	step4 := utils.SHA256(step3)
	step5 := utils.SHA256(step4)
	step6 := step5[:4]
	step7 := append(step3, step6...)
	return utils.EncodeToBase58(utils.EncodeRipple, step7)
}
