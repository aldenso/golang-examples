/*
golang example to tell goroutines to stop after you found the value you were
searching in a large wordlist file.
*/
package goroutineschannelquit

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var initTime = time.Now()

var wordlistfile string = "all.lst"

var pattern string = "pattern"

// Define fileScanner with methods
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

// Define a worker job with methods
type Worker struct {
	quit chan bool
}

func NewWorker() *Worker {
	f := &Worker{make(chan bool)}
	go f.run()
	return f
}

func (f *Worker) run() {
	for {
		select {
		case <-f.quit:
			fmt.Println("finishing task")
			time.Sleep(time.Second)
			fmt.Println("task done")
			f.quit <- true
			return
		case <-time.After(time.Second):
			fmt.Println("still running task")
		}
	}
}

func (f *Worker) Stop() {
	fmt.Println("server stopping")
	f.quit <- true
	<-f.quit
	fmt.Println("server stopped")
}

func goroutineschannelquit() {
	runner := NewWorker()
	fscanner := NewFileScanner()
	err := fscanner.Open(wordlistfile)
	if err == nil {
		scanner := fscanner.GetScan()
		for scanner.Scan() {
			value := scanner.Text()
			if value == pattern {
				end := time.Now()
				d := end.Sub(initTime)
				duration := d.Seconds()
				fmt.Printf("pattern found: %s\n", value)
				fmt.Printf("completed in %v senconds\n", strconv.FormatFloat(duration, 'g', -1, 64))
				runner.Stop()
			}
		}
		fscanner.Close()
		//time.Sleep(1 * time.Second)
	}
}
