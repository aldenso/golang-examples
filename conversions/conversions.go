package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = "this is a string"

func main() {
	// convert string to slice of strings
	slice := strings.Split(input, " ")
	// replace all the space caracthers with comma in a string
	replaced := strings.Replace(input, " ", ",", -1)
	// show only first string in a slice of strings
	firstStringInSlice := string(slice[0])
	// show every string in a slice except for the first "0"
	nonFirstStringInSlice := slice[1:]
	// convert every element in slice to a slice of quoted strings
	quotedSlice := []string{}
	for _, i := range slice {
		quotedSlice = append(quotedSlice, strconv.Quote(i))
	}
	// convert string to bytes slice
	stringToBytes := []byte(input)
	// convert bytes slice to a string
	bytesToString := string(stringToBytes)
	fmt.Println("input:", input)
	fmt.Println("slice:", slice)
	fmt.Println("replaced:", replaced)
	fmt.Println("firstStringInSlice:", firstStringInSlice)
	fmt.Println("nonFirstStringInSlice:", nonFirstStringInSlice)
	fmt.Println("quotedSlice:", quotedSlice)
	fmt.Println("stringToBytes:", stringToBytes)
	fmt.Println("bytesToString:", bytesToString)
}
