package repl

import(
    "bufio"
    "fmt"
    "io"
    "monkey/lexer"
    "monkey/token"
)

const PROMPT = ">>"

func Start(in io.Reader) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Printf(PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return
        }

        line := scanner.Text()

        l := lexer.New(line)

        tok := l.NextToken()
        for tok.Type != token.EOF {
            fmt.Printf("%+v\n",tok)
            tok = l.NextToken()
        }
    }
}
