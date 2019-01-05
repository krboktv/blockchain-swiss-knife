package main

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"bytes"
	"encoding/hex"
	"fmt"
	"encoding/json"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

type Transaction struct {
TxId               string `json:"txid"`
SourceAddress      string `json:"source_address"`
DestinationAddress string `json:"destination_address"`
Amount             int64  `json:"amount"`
UnsignedTx         string `json:"unsignedtx"`
SignedTx           string `json:"signedtx"`
}

func CreateTransaction(secret string, destination string, amount int64, txHash string) (Transaction, error) {
var transaction Transaction
wif, err := btcutil.DecodeWIF(secret)
if err != nil {
return Transaction{}, err
}
addresspubkey, _ := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), &chaincfg.MainNetParams)
sourceTx := wire.NewMsgTx(wire.TxVersion)
sourceUtxoHash, _ := chainhash.NewHashFromStr(txHash)
sourceUtxo := wire.NewOutPoint(sourceUtxoHash, 0)
sourceTxIn := wire.NewTxIn(sourceUtxo, nil, nil)
destinationAddress, err := btcutil.DecodeAddress(destination, &chaincfg.MainNetParams)
sourceAddress, err := btcutil.DecodeAddress(addresspubkey.EncodeAddress(), &chaincfg.MainNetParams)
if err != nil {
return Transaction{}, err
}
destinationPkScript, _ := txscript.PayToAddrScript(destinationAddress)
sourcePkScript, _ := txscript.PayToAddrScript(sourceAddress)
sourceTxOut := wire.NewTxOut(amount, sourcePkScript)
sourceTx.AddTxIn(sourceTxIn)
sourceTx.AddTxOut(sourceTxOut)
sourceTxHash := sourceTx.TxHash()
redeemTx := wire.NewMsgTx(wire.TxVersion)
prevOut := wire.NewOutPoint(&sourceTxHash, 0)
redeemTxIn := wire.NewTxIn(prevOut, nil, nil)
redeemTx.AddTxIn(redeemTxIn)
redeemTxOut := wire.NewTxOut(amount, destinationPkScript)
redeemTx.AddTxOut(redeemTxOut)
sigScript, err := txscript.SignatureScript(redeemTx, 0, sourceTx.TxOut[0].PkScript, txscript.SigHashAll, wif.PrivKey, false)
if err != nil {
return Transaction{}, err
}
redeemTx.TxIn[0].SignatureScript = sigScript
flags := txscript.StandardVerifyFlags
vm, err := txscript.NewEngine(sourceTx.TxOut[0].PkScript, redeemTx, 0, flags, nil, nil, amount)
if err != nil {
return Transaction{}, err
}
if err := vm.Execute(); err != nil {
return Transaction{}, err
}
var unsignedTx bytes.Buffer
var signedTx bytes.Buffer
sourceTx.Serialize(&unsignedTx)
redeemTx.Serialize(&signedTx)
transaction.TxId = sourceTxHash.String()
transaction.UnsignedTx = hex.EncodeToString(unsignedTx.Bytes())
transaction.Amount = amount
transaction.SignedTx = hex.EncodeToString(signedTx.Bytes())
transaction.SourceAddress = sourceAddress.EncodeAddress()
transaction.DestinationAddress = destinationAddress.EncodeAddress()
return transaction, nil
}

func main(){
	transaction, err := CreateTransaction("5HusYj2b2x4nroApgfvaSfKYZhRbKFH41bVyPooymbC6KfgSXdD", "1KKKK6N21XKo48zWKuQKXdvSsCf95ibHFa", 91234, "81b4c832d70cb56ff957589752eb4125a4cab78a25a8fc52d6a09e5bd4404d48")
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := json.Marshal(transaction)
	fmt.Println(string(data))

}
