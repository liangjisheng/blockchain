package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/blockcypher/gobcy"
)

const (
	// Token my blockcypher token
	Token = "61a2e20f1df54fe5a0ad9226ce0d8b21"
)

func getTx(txHash, coin, chain string) string {
	bcy := gobcy.API{
		Token: Token,
		Coin:  coin,
		Chain: chain,
	}
	tx, err := bcy.GetTX(txHash, nil)
	if err != nil {
		fmt.Println("error:", err)
	}
	return fmt.Sprintf("%+v", tx)
}

func getAddrBalance(addr, coin, chain string) string {
	bcy := gobcy.API{
		Token: Token,
		Coin:  coin,
		Chain: chain,
	}
	bal, err := bcy.GetAddrBal(addr, nil)
	if err != nil {
		fmt.Println("error:", err)
	}
	return fmt.Sprintf("%+v", bal)
}

func getAddr(coin, chain string) string {
	bcy := gobcy.API{
		Token: Token,
		Coin:  coin,
		Chain: chain,
	}
	addrKeys, err := bcy.GenAddrKeychain()
	if err != nil {
		fmt.Println("error:", err)
	}
	return fmt.Sprintf("%+v", addrKeys)
}

func broadcastTx(txHex, coin, chain string) string {
	bcy := gobcy.API{
		Token: Token,
		Coin:  coin,
		Chain: chain,
	}
	skel, err := bcy.PushTX(txHex)
	if err != nil {
		fmt.Println("error:", err)
	}
	return fmt.Sprintf("%+v", skel)
}

func getAddrBalance1(addr string) {
	// GET https://chain.so/api/v2/get_address_balance/{NETWORK}/{ADDRESS}[/{MINIMUM CONFIRMATIONS}]
	url := "https://chain.so/api/v2/get_address_balance/ltc/" + addr
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// fmt.Printf("%+v", *resp)
	fmt.Println(string(body))
}

func broadcastTx1(txHex, coin string) {
	// curl -d 'tx_hex=0102100001ac…' https://chain.so/api/v2/send_tx/DOGE
	var data bytes.Buffer
	tx := struct {
		TxHex string `json:"tx_hex"`
	}{
		TxHex: txHex,
	}
	enc := json.NewEncoder(&data)
	if err := enc.Encode(tx); err != nil {
		return
	}
	fmt.Println(data.String())

	url := "https://chain.so/api/v2/send_tx/" + coin
	resp, err := http.Post(url, "application/json", &data)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		fmt.Println("status code:", resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
		return
	}

	dec := json.NewDecoder(resp.Body)
	var decTarget interface{}
	err = dec.Decode(decTarget)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%+v\n", decTarget)
}

func main() {
	// txHash := "3094c7e15d635e440987842103968b1782274341b6330154e52f858cf6d500be"
	// fmt.Println(getTx(txHash, "ltc", "main"))

	// addr := "LSjMSjutEpn356z5mpF5yx9PPLoSmuGAuz"
	// fmt.Println(getAddrBalance(addr, "ltc", "main"))
	// fmt.Println(getAddr("ltc", "main"))

	txHex := "12345667847472475847423843484747348337373747848345"
	// fmt.Println(broadcastTx(txHex, "ltc", "main"))

	// getAddrBalance1(addr)
	broadcastTx1(txHex, "ltc")
}

// curl -d '{"tx":""}' https://api.blockcypher.com/v1/ltc/main/txs/push?token=61a2e20f1df54fe5a0ad9226ce0d8b21
// curl -d '{"tx_hex":""}' https://chain.so/api/v2/send_tx/ltc
