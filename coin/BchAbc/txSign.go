package main

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/bcext/gcash/btcec"
	"github.com/bcext/gcash/txscript"
	"github.com/bcext/gcash/wire"
)

func txSign() {
	rawPriv, err := hex.DecodeString("3957ae9390a1341738a8fb667e0fb9e0ffd01ce0e4ee4de1b3a6deb162c81b48")
	if err != nil {
		panic(err)
	}
	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), rawPriv)
	rawTx, err := hex.DecodeString("0200000001ceda94b0beabb9510bc03e803c8aaceef736ea40d5d8a4d943927d1135cd57e40100000000ffffffff0300e1f505000000001976a9140000000000000000000000000000000000376e4388ac00000000000000000a6a080877686300000044e00f9700000000001976a914f9757e1b9caa87ef4377d977754ae95908cc6df188ac00000000")
	if err != nil {
		panic(err)
	}
	// wif, err := cashutil.DecodeWIF("cSnfjz72jSMQhdxQeQ6uW8eUko9xndFoZ5w5fTvbSCNgeTjaUgQW")
	// if err != nil {
	// 	panic(err)
	// }
	// use:     wif.PrivKey.PubKey()

	var tx wire.MsgTx
	err = tx.Deserialize(bytes.NewReader(rawTx))
	if err != nil {
		panic(err)
	}
	scriptPubKey, err := hex.DecodeString("76a914f9757e1b9caa87ef4377d977754ae95908cc6df188ac")
	if err != nil {
		panic(err)
	}
	sig, err := txscript.RawTxInSignature(&tx, 0, scriptPubKey, 110000000, txscript.SigHashAll|txscript.SigHashForkID, privKey)
	if err != nil {
		panic(err)
	}
	sig, err = txscript.NewScriptBuilder().AddData(sig).Script()
	if err != nil {
		panic(err)
	}
	pk, err := txscript.NewScriptBuilder().AddData(pubKey.SerializeCompressed()).Script()
	if err != nil {
		panic(err)
	}
	sig = append(sig, pk...)
	tx.TxIn[0].SignatureScript = sig
	fmt.Println("signature", hex.EncodeToString(sig))
	w := bytes.NewBuffer(make([]byte, 0))
	err = tx.Serialize(w)
	if err != nil {
		panic(err)
	}
	fmt.Println("transaction", hex.EncodeToString(w.Bytes()))

	engine, err := txscript.NewEngine(scriptPubKey, &tx, 0, txscript.StandardVerifyFlags, nil, nil, 110000000)
	if err != nil {
		panic(err)
	}

	// verify the signature
	err = engine.Execute()
	if err != nil {
		panic(err)
	}
}
