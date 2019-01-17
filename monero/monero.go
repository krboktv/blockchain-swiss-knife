package monero

type Monero struct {
	Seed            string
	PrivateKeyVeiw  string
	PrivateKeySpend string
	PublicKeyVeiw   string
	PublicKeySpend  string
	Address         string
}

var (
	MainnetXMR = []byte{0x12}
)

//func generateRandomSeed() ([]byte, error) {
//	passphrase, err := generateRandomPassphrase()
//	if err != nil {
//		return nil, err
//	}
//	seed := utils.SHA512(passphrase)[:32]
//	encodedSeed, err := Encode(VersionByteSeed, seed)
//	return encodedSeed, err
//}
