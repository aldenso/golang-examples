package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var input = "this is a string"
var numstring = "-50"
var timeInput = time.Now()

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
	// convert time format to RFC3339
	timeInputToRFC3339 := timeInput.Format(time.RFC3339)
	// convert to unix time format
	timeInputToUnix := timeInput.Unix()
	// not really converting, but get the timeframe between timeInput and this timePeriodSeconds execution
	timePeriodSeconds := time.Since(timeInput).Seconds()
	// convert string to int
	numstringtoint, _ := strconv.Atoi(numstring)

	fmt.Println("input:", input, "|| type:", reflect.TypeOf(input))
	fmt.Println("slice:", slice, "|| type:", reflect.TypeOf(slice))
	fmt.Println("replaced:", replaced)
	fmt.Println("firstStringInSlice:", firstStringInSlice)
	fmt.Println("nonFirstStringInSlice:", nonFirstStringInSlice)
	fmt.Println("quotedSlice:", quotedSlice)
	fmt.Println("stringToBytes:", stringToBytes)
	fmt.Println("bytesToString:", bytesToString)
	fmt.Println("timeInput:", timeInput)
	fmt.Println("timeInputToRFC3339:", timeInputToRFC3339)
	fmt.Println("timeInputToUnix:", timeInputToUnix)
	fmt.Println("timePeriodSeconds:", timePeriodSeconds)
	fmt.Println("numstring:", numstring, "|| type:", reflect.TypeOf(numstring))
	fmt.Println("numstringtoint:", numstringtoint, "|| type:", reflect.TypeOf(numstringtoint))
}
