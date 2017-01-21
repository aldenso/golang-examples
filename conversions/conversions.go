/*
Golang examples for my most regular conversions.
*/
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

const longForm = "Jan 2, 2006 at 3:04pm (MST)"

var timeInput, _ = time.Parse(longForm, "Jan 21, 2017 at 3:30am (VET)")
var time4since, _ = time.Parse(longForm, "Jan 22, 2017 at 2:30am (VET)")

func runGeneralConversions() {
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
	fmt.Println("numstring:", numstring, "|| type:", reflect.TypeOf(numstring))
	fmt.Println("numstringtoint:", numstringtoint, "|| type:", reflect.TypeOf(numstringtoint))
}

func runTimeConversions() {
	// convert time format to RFC3339
	timeInputToRFC3339 := timeInput.Format(time.RFC3339)
	// convert to unix time format
	timeInputToUnix := timeInput.Unix()
	fmt.Println("timeInput:", timeInput)
	fmt.Println("timeInputToRFC3339:", timeInputToRFC3339)
	fmt.Println("timeInputToUnix:", timeInputToUnix)
}

func runNotRealConversion() {
	// not really converting, but get the timeframe between timeInput and this timePeriodSeconds execution
	timePeriodSeconds := time.Since(timeInput).Seconds()
	fmt.Println("timePeriodSeconds:", timePeriodSeconds)
}

func main() {
	runGeneralConversions()
	runTimeConversions()
	runNotRealConversion()
}
