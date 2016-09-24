package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Command struct for commands, arguments, stdout and stderr.
type Command struct {
	Cmd    string
	Args   []string
	Output []byte
	Err    error
}

var cmd1, cmd2, cmd3 Command

func main() {
	home := os.Getenv("HOME")
	os.Setenv("PATH", "/usr/bin:/sbin")
	cmd1.Cmd, cmd1.Args = "ls", []string{"-la", home}
	cmd2.Cmd, cmd2.Args = "date", []string{""}
	cmd3.Cmd, cmd3.Args = "df", []string{"-h"}
	var cmds = []Command{cmd1, cmd2, cmd3}
	for _, item := range cmds {
		fmt.Printf("Running Command: %s, with args: %s \n", item.Cmd, item.Args)
		item.Output, item.Err = exec.Command(item.Cmd, item.Args...).Output()
		if item.Err != nil {
			fmt.Println(item.Err)
		}
		fmt.Println(string(item.Output))
	}
}
