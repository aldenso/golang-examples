/*
Golang examples to show a terminal output using a pager.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func pager(input string) {
	pager := os.Getenv("PAGER")
	if pager == "" {
		pager = "more"
	}
	cmd := exec.Command(pager)
	cmd.Stdin = strings.NewReader(input)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	wantedOutput := fmt.Sprintf(`
just
about
to
start
a
very
long
output
to
show
%s
the
use
of
a
pager
in
a
regular
terminal
%s
plus
some
more
lines
because
I
want
to`, "you", "size")
	pager(wantedOutput)
}
