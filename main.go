package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	testRequest := `{"jsonrpc":"1.0","id":"curltext","method":"getblockchaininfo","params":[]}`
	url := fmt.Sprintf("http://%s:%s@192.168.1.154:8332/", os.Getenv("RPCUSER"), os.Getenv("RPCPASSWORD"))
	req, _ := http.NewRequest("POST", url, strings.NewReader(testRequest))
	req.Header.Add("content-type", "text/plain;")

	res, e := http.DefaultClient.Do(req)
	if e != nil {
		fmt.Println(e)
	} else {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
}
