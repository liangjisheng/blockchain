package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func broadTx(rawHex string) {
	var data bytes.Buffer
	rawTx := struct {
		RawHex string `json:"rawtx"`
	}{
		RawHex: rawHex,
	}

	enc := json.NewEncoder(&data)
	if err := enc.Encode(rawTx); err != nil {
		fmt.Println("encode error:", err)
		return
	}
	fmt.Println(data.String())

	// https://developers.whatsonchain.com
	url := "https://bchsvexplorer.com/api/tx/send/"
	resp, err := http.Post(url, "application/json", &data)
	if err != nil {
		fmt.Println("post error:", err)
		return
	}
	defer resp.Body.Close()
	// fmt.Println("statusCode:", resp.StatusCode)
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
}

func main() {
	var rawHex string
	flag.StringVar(&rawHex, "rawHex", "", "input raw hex tx.")
	flag.Parse()

	// rawHex = "0100000001656683dc5ab95dab3ce4bf2ff22be783a3524a456ef84953a7ac65ab72a74cbe000000006b483045022100b3d7c2d27cedc9e3cada555b6e3ba841399a5b2e096d52b87faac7bf1a5a200b0220527d74b0f70b374b1d38bd66a929443719b44bb45af89db45c79936e9225d3174121037eb1f69162d22c703884dba9b72b9447adff289ac2c75500c0794c57c6b5a589ffffff00019b5fcd1d000000001976a9144cd6b1724d140d0e19e23bfbfac325d8723072a388ac00000000"
	broadTx(rawHex)
}
