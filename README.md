# blockchain-swiss-knife

| Method | Input | Output |
| ------ | ------ | ------ |
| GenerateKey | nil | privateKey []byte, error |
| GetPublicKey | privateKey []byte | publicKey []byte |
| GetAddress | privateKey []byte | address []byte |
| GetBalance | address string | balance float64, error |
| SignTransaction | privateKey []byte, receiver []byte, amount uint | rawTransaction []byte, error |
| SendTransaction | rawTransaction []byte | transactionHash string, error |
