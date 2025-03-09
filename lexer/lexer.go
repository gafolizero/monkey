package lexer

import "monkey/token"

type Lexer struct {
    input string
    position int // current ch index
    readPosition int // next ch index
    ch byte
}

func New(input string) *Lexer {
    l := &Lexer{ input: input }
    l.readChar()
    return l
}

func (l *Lexer) readChar() {
    if (*l).readPosition >= len((*l).input) {
        (*l).ch = 0
    }else {
        (*l).ch = (*l).input[(*l).readPosition]
    }

    (*l).position = (*l).readPosition
    (*l).readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type : tokenType, Literal: string(ch)}
}

//type Token struct {
//    Type TokenType
//    Literal string
//}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
    position := (*l).position
    for isLetter((*l).ch) {
        l.readChar()
    }
    return (*l).input[position:(*l).position]
}

func isDigit(ch byte) bool{
    return ch >= '0' && ch <= '9'
}

func (l *Lexer) readDigit() string{
    position := (*l).position
    for isDigit((*l).ch) {
        l.readChar()
    }
    return (*l).input[position: (*l).position]
}

func (l *Lexer) ignoreWhiteSpace() {
    for (*l).ch == ' ' || (*l).ch == '\t' || (*l).ch == '\n' || (*l).ch == '\r' {
       l.readChar()
    }
}

func (l *Lexer) peekChar() byte {
    if (*l).readPosition >= len((*l).input) {
        return 0
    }
    return (*l).input[(*l).readPosition]
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token
    l.ignoreWhiteSpace()
    switch (*l).ch {
    case '=':
        peek := l.peekChar()
        if peek == '='{
            tok.Literal = string((*l).ch) + string(peek)
            tok.Type = token.EQ
            l.readChar()
            l.readChar()
            return tok
        }
        tok = newToken(token.ASSIGN, (*l).ch)
    case ';':
        tok = newToken(token.SEMICOLON, (*l).ch)
    case '(':
        tok = newToken(token.LPAREN, (*l).ch)
    case ')':
        tok = newToken(token.RPAREN, (*l).ch)
    case '{':
        tok = newToken(token.LBRACE, (*l).ch)
    case '}':
        tok = newToken(token.RBRACE, (*l).ch)
    case ',':
        tok = newToken(token.COMMA, (*l).ch)
    case '+':
        tok = newToken(token.PLUS, (*l).ch)
    case '-':
        tok = newToken(token.MINUS, (*l).ch)
    case '!':
        peek := l.peekChar()
        if peek == '='{
            tok.Literal = string((*l).ch) + string(peek)
            tok.Type = token.NOT_EQ
            l.readChar()
            l.readChar()
            return tok
        }
        tok = newToken(token.BANG, (*l).ch)
    case '*':
        tok = newToken(token.ASTERISK, (*l).ch)
    case '/':
        tok = newToken(token.SLASH, (*l).ch)
    case '<':
        tok = newToken(token.LT, (*l).ch)
    case '>':
        tok = newToken(token.GT, (*l).ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    default:
        if isLetter((*l).ch){
            tok.Literal = l.readIdentifier()
            tok.Type = token.LookupIdent(tok.Literal)
            return tok
        }else if isDigit((*l).ch){
            tok.Literal = l.readDigit()
            tok.Type = token.INT
            return tok
        }else {
            tok = newToken(token.ILLEGAL, (*l).ch)
        }
    }
    l.readChar()
    return tok
}
