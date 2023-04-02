package lexer

import "token"

type Lexer struct {
  input        string
  position     int    // current position
  readPosition int    // next position
  ch           byte   // current character
}

func New(input string) *Lexer {
  l := &Lexer(input: input)
  l.readChar()
  return l
}

func (l *Lexer) readChar() {
  if l.readPosition >= len(l.input) {
    l.ch = 0
  } else {
    l.ch = l.input[l.readPosition]
  }
  l.position = l.readPosition
  l.readPosition += 1
}

func (l * Lexer) peekChar() byte {
  if l.readPosition >= len(l.input) {
    return 0
  } else {
    return l.input[l.readPosition]
  }
}

func (i *Lexer) NextToken() token.Token {
  var tok token.Token

  l.skipWhitespace()

  switch l.ch {
  case ';':
    tok = newToken(token.SEMICOLON, l.ch)
  case 0:
    tok.Literal = ""
    tok.Type = token.EOF
  default:
    if isDigit(l.ch) {
      tok.Type = token.NUM
      tok.Literal = l.readNumber()
      return tok
    } else {
      tok = newToken(token.ILLEGAL, l.ch)
    }
  }

  l.readChar()
  return tok
}

func newToken(TokenType token.TokenType, ch byte) token.Token {
  return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z'|| 'A' < ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
  position := l.position
  for isDigit(l.ch) {
    l.readChar()
  }
  return l.input[position:l.position]
}

func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
  for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
    l.readChar()
  }
}
