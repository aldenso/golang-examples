/*
Small golang example for creation of struct and methods to read a file
*/
package filescannerstruct

import (
	"bufio"
	"fmt"
	"os"
)

var wordlistfile string = "all.lst"

var pattern string = "pattern"

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

func filescannerstruct() {
	fscanner := NewFileScanner()
	err := fscanner.Open(wordlistfile)
	if err == nil {
		scanner := fscanner.GetScan()
		for scanner.Scan() {
			line := scanner.Text()
			if line == pattern {
				fmt.Println(line)
				fscanner.Close()
			}
		}
	}
}
