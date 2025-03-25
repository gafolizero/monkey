package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`
	fmt.Print(MONKEY_FACE)
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s\n", user.Username)
	input := os.Stdin
	repl.Start(input, os.Stdout)
}
