package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func doc1(rawHex, doMainName, method string) {
	var data bytes.Buffer
	rawTx := struct {
		RawHex string `json:"rawhex"`
	}{
		RawHex: rawHex,
	}

	enc := json.NewEncoder(&data)
	if err := enc.Encode(rawTx); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(data.String())

	url := doMainName + method
	fmt.Println(url)
	resp, err := http.Post(url, "application/json", &data)
	if err != nil {
		fmt.Println("post error:", err)
		return
	}
	defer resp.Body.Close()
	// fmt.Println("status code:", resp.StatusCode)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		fmt.Println("status code:", resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
		return
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read all error:", err)
		return
	}
	fmt.Println(string(res))

	// dec := json.NewDecoder(resp.Body)
	// var decTarget interface{}
	// err = dec.Decode(decTarget)
	// if err != nil {
	// 	fmt.Println("json decode error:", err)
	// 	return
	// }
	// fmt.Printf("%+v\n", decTarget)
}

func doc2(rawHex, pattern string) {
	var data bytes.Buffer
	rawTx := struct {
		RawHex []string `json:"hexes"`
	}{
		RawHex: []string{rawHex},
	}

	enc := json.NewEncoder(&data)
	if err := enc.Encode(rawTx); err != nil {
		fmt.Println("error:", err)
		return
	}
	// fmt.Println(data.String())

	url := "https://rest.bitcoin.com/" + pattern
	resp, err := http.Post(url, "application/json", &data)
	if err != nil {
		fmt.Println("post error:", err)
		return
	}
	defer resp.Body.Close()
	// fmt.Println("status code:", resp.StatusCode)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		fmt.Println("status code:", resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
		return
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read all error:", err)
		return
	}
	fmt.Println(string(res))

	// dec := json.NewDecoder(resp.Body)
	// var decTarget interface{}
	// err = dec.Decode(decTarget)
	// if err != nil {
	// 	fmt.Println("json decode error:", err)
	// 	return
	// }
	// fmt.Printf("%+v\n", decTarget)
}

func main() {
	rawHex := "01000000012a00f2004283d110630f5f4d06fc1b5d7c36305f842741968d4636b08e54cfef000000006a47304402200e1ddbe04503ab14c89579200ff566c33cd6225375bb2cc0d44c2d25749b6c8a022077ad6a480e48cd8c755a86a24eb124a2ac2371f612052e7ae7304c41cc97b1a7412103ccd9e0e3cdf92d05955dcbe2027ce198cbb1744d65dc6a4a9accb3c6831a22ecffffffff02a0860100000000001976a9145fa859d3f4de61d8dc2e9b36e7da23f86d126e1488ace8af0d00000000001976a914308c6ba1440434aef594e6db9e549e8fdac237fa88ac00000000"

	// doc1(rawHex, "https://bch-chain.api.btc.com/", "v3/tools/tx-decode")
	// doc1(rawHex, "https://bch-chain.api.btc.com/", "v3/tools/tx-publish")

	// doc1(rawHex, "https://bch.btc.com/service/", "tx/decode")
	// doc1(rawHex, "https://bch.btc.com/service/", "tx/publish")

	// https://developer.bitcoin.com/rest/docs/rawtransactions
	// doc2(rawHex, "v2/rawtransactions/decodeRawTransaction")
	doc2(rawHex, "v2/rawtransactions/sendRawTransaction")

	// txSign()
}
