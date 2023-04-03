package token

type TokenType string

type Token struct {
  Type    TokenType
  Literal string
}

const (
  ILLEGAL   = "ILLEGAL"
  EOF       = "EOF"

  // identifier + literal
  IDENT     = "IDENT"   // int, float, a, b, ...
  NUM       = "NUM"     // 1, 2, 3,...

  // operator
  ASSIGN    = "="
  PLUS      = "+"
  MINUS     = "-"
  BANG      = "!"
  ASTERISK  = "*"
  SLASH     = "/"

  LT        = "<"
  GT        = ">"

  EQ        = "=="
  NOT_EQ    = "!="

  // delimiter
  COMMA     = ","
  SEMICOLON = ";"

  LPAREN    = "("
  RPAREN    = ")"
  LBRACE    = "{"
  RBRACE    = "}"

  // keywords
  INT       = "int"
  FLOAT     = "float"
)

var keywords = map[string]TokenType{
  "int":       INT,
  "float":     FLOAT,
}

func LookupIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}
