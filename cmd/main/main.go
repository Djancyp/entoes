package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	test "net/url"
	"os"
	"strings"
)

type Response struct {
	Text []string `json:"text"`
}

type PrintRespose struct {
	Text []string `json:"text"`
}
type Payload struct {
	Text    string `json:"text"`
	Options int    `json:"options"`
}
type Values map[string][]string

func main() {
	// read from arg 1
	word := os.Args[1]
	if word == "" {
		log.Fatalln("No word provided")
	}
	url := "https://translate.yandex.net/api/v1/tr.json/translate?id=7f451d1e.630de6b7.1270c0f4.74722d74657874-11-0&srv=tr-text&lang=en-es&reason=auto&format=text&ajax=1&yu=5342753661661855320"
	data := test.Values{}
	data.Set("text", word)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var printResponse PrintRespose
	err = json.Unmarshal(body, &printResponse)
	if err != nil {
		fmt.Println("This is the err", err)
		panic(err)
	}
	fmt.Println(printResponse.Text[0])
}
