package lexer

import (
  "testing"

  "token"
)

func TestNextToken(t *testing.T) {
  input := `int a;
a = 42;
return 0;
`

  tests := []struct {
    expectedType    token.TokenType
    expectedLiteral string
  }{
    {token.INT, "int"},
    {token.IDENT, "a"},
    {token.SEMICOLON, ";"},
    {token.IDENT, "a"},
    {token.ASSIGN, "="},
    {token.NUM, "42"},
    {token.SEMICOLON, ";"},
    {token.RETURN, "return"},
    {token.NUM, "0"},
    {token.SEMICOLON, ";"},
  }

  l := New(input)

  for i, tt := range tests {
    tok := l.NextToken()

    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
        i, tt.expectedType, tok.Type)
    }

    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
        i, tt.expectedLiteral, tok.Literal)
    }
  }
}
