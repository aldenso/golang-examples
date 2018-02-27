package main

import (
	"fmt"
	"log"
	"net/http"
)

// getURL simple function to make http client requests.
func getURL(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res, nil
}

// getUserInput simple function to get user input from stdin.
func getUserInput() error {
	fmt.Println("Getting user input")
	var answer string
	fmt.Println("input (N|n - Y|y - anykey):")
	_, err := fmt.Scanf("%s\n", &answer)
	if err != nil {
		return err
	}
	switch {
	case answer == "N" || answer == "n":
		fmt.Printf("Got case for 'N' or 'n' - '%s'\n", answer)
		return nil
	case answer == "Y" || answer == "y":
		fmt.Printf("Got case for 'Y' or 'y' - '%s'\n", answer)
		return nil
	default:
		fmt.Printf("Got default case for '%s'\n", answer)
		return nil
	}
}

func main() {
	url := "https://www.google.com"
	// discard response
	_, err := getURL(url)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	fmt.Println("data was retrieved")

	err = getUserInput()
	if err != nil {
		log.Fatalln("Error getting user input: ", err)
	}
}
