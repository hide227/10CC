package token

type TokenType string

type Token struct {
  Type    TokenType
  Literal string
}

const (
  ILLEGAL   = "ILLEGAL"
  EOF       = "EOF"

  NUM       = "NUM"     // 1, 2, 3,...

  // operator
  ASSIGN    = "="
  PLUS      = "+"

  // delimiter
  COMMA     = ","
  SEMICOLON = ";"

  LPAREN    = "("
  RPAREN    = ")"
  LBRACE    = "{"
  RBRACE    = "}"
)
