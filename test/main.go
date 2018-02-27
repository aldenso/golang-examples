package main

import (
	"fmt"
	"log"
	"net/http"
)

func getURL(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res, nil
}

func main() {
	url := "https://www.google.com"
	// discard response
	_, err := getURL(url)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	fmt.Println("data was retrieved")
}
