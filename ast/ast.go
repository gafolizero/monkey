package ast

import (
    "monkey/token"
    "bytes"
)

type Node interface {
    TokenLiteral() string
    String() string
}

type Statement interface {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}

// ROOT Node
type Program struct {
    Statements []Statement
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}

func (p *Program) String() string {
    var out bytes.Buffer

    for _, s := range p.Statements {
        out.WriteString(s.String())
    }

    return out.String()
}

// Node for let statement
type LetStatement struct {
    Token token.Token
    Name *Identifier
    Value Expression
}

// Token { Type : "IDENT", Literal: "foo" }

func (ls *LetStatement) TokenLiteral () string {
    return ls.Token.Literal
}

func (ls *LetStatement) statementNode () {}

func (ls *LetStatement) String() string {
    var out bytes.Buffer

    out.WriteString(ls.TokenLiteral() + " ")
    out.WriteString(ls.Name.String())
    out.WriteString(" = ")

    if ls.Value != nil {
        out.WriteString(ls.Value.String())
    }
    out.WriteString(";")

    return out.String()
}

// Node for identifier
type Identifier struct {
    Token token.Token
    Value string
}

func (i *Identifier) expressionNode () {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String () string { return i.Value }

// Node for return statement
type ReturnStatement struct {
    Token token.Token
    ReturnValue Expression
}

func (rs *ReturnStatement) statementNode(){}
func (rs *ReturnStatement) TokenLiteral() string {return rs.Token.Literal}

func (rs *ReturnStatement) String() string {
    var out bytes.Buffer

    out.WriteString(rs.TokenLiteral() + " ")
    if rs.ReturnValue != nil {
        out.WriteString(rs.ReturnValue.String())
    }
    return out.String()
}

// Node for return statement
type ExpressionStatement struct {
    Token token.Token
    Expression Expression
}

func (rs *ExpressionStatement) statementNode() {}
func (rs *ExpressionStatement) TokenLiteral() string {return rs.Token.Literal}

func (es *ExpressionStatement) String() string {
    if es.Expression != nil {
        return es.Expression.String()
    }
    return ""
}


