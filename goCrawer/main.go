package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	resp, err := http.Get("https://golang.google.cn/pkg/")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Printf(bodyString)
	}
}
