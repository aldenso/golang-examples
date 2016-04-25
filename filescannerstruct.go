/*
Small golang example for creation of struct and methods to read a file
*/
package filescanner

import (
	"bufio"
	"fmt"
	"os"
)

var passwordfile string = "passwords.txt"

/*
struct for passwords file scan, with methods
fileScanner --> struct
NewFileScanner --> Object
Open --> method
Close --> method
GetScan --> method
*/
type fileScanner struct {
	File    *os.File
	Scanner *bufio.Scanner
}

func NewFileScanner() *fileScanner {
	return &fileScanner{}
}

func (f *fileScanner) Open(path string) (err error) {
	f.File, err = os.Open(path)
	return err
}

func (f *fileScanner) Close() error {
	return f.File.Close()
}

func (f *fileScanner) GetScan() *bufio.Scanner {
	if f.Scanner == nil {
		f.Scanner = bufio.NewScanner(f.File)
		f.Scanner.Split(bufio.ScanLines)
	}
	return f.Scanner
}

func filescanner() {
	fscanner := NewFileScanner()
	err := fscanner.Open(passwordfile)
	if err == nil {
		scanner := fscanner.GetScan()
		for scanner.Scan() {
			line := scanner.Text()
			if line == "Okm00okm" {
				fmt.Println(line)
				fscanner.Close()
			}
		}
	}
}
