package ast

import (
	"monkey/token"
	"testing"
)

/*
type LetStatement struct {
    Token token.Token
    Name *Identifier
    Value Expression
}
*/
// let myVar = anotherVar;
func TestString(t *testing.T) {
	program := &Program{
		[]Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong, got=%q", program.String())
	}
}
