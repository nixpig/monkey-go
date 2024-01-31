package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/nixpig/monkey-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey REPL!\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
