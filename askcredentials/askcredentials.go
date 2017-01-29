package main

import (
	"fmt"
	"log"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func credentials() (string, string) {
	fmt.Printf("Enter Username:\n")
	var username string
	_, err := fmt.Scanln(&username)
	if err != nil {
		log.Fatalf("Failed to read Username: '%v'\n", err)
	}

	fmt.Printf("Enter Password:\n")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalf("Failed to read Password: '%v'\n", err)
	}
	password := string(bytePassword)

	return username, strings.TrimSpace(password)
}
func main() {
	user, pass := credentials()
	fmt.Printf("Username: %s Password: %s\n", user, pass)
}
