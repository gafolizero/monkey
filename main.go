package main

import (
    "fmt"
    "os"
    "os/user"
    "monkey/repl"
)

func main() {
    user, err := user.Current()

    if err != nil {
        panic(err)
    }

    fmt.Printf("Hello %s\n", user.Username)
    input := os.Stdin
    repl.Start(input)
}
