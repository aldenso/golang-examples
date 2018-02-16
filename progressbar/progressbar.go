//Just a progressbar example.
package main

import (
	"fmt"
	"time"

	"github.com/schollz/progressbar"
)

func main() {
	maxnumber := 10
	bar := progressbar.New(maxnumber)
	bar.SetTheme([]string{"#", "=", "|", "|"})

	fmt.Println("Run first loop: not complete")
	for x := 0; x < maxnumber; x++ {
		bar.Add(1)
		time.Sleep(500 * time.Millisecond)
		if x == 5 {
			break
		}
	}
	fmt.Println("Break.")

	fmt.Println("Run second loop: complete")
	bar.Reset()
	for x := 0; x < maxnumber; x++ {
		bar.Add(1)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Done.")
}
