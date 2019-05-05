package main

import (
	"fmt"
	"os"
	user2 "os/user"
	"repl"
)

func main() {
	user, err := user2.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Zdravo %s! Ovo je programski jezik Monkey!\n", user.Username)
	fmt.Printf("Otvoren je REPL za izvrsavanje komandi\n")

	repl.Start(os.Stdin, os.Stdout)
}
