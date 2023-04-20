package ast

import "token"

type Node interface {
  TokenLiteral() string
}

type Statement interface {
  Node
  statementNode()
}

type Expression interface {
  Node
  expressionNode()
}

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

type IntStatement struct {
  Token token.Token // token.INT token
  Name *Identifier
  Value Expression
}

func (is *IntStatement) statementNode()       {}
func (is *IntStatement) TokenLiteral() string { return is.Token.Literal }

type Identifier struct {
  Token token.Token // token.IDENT token
  Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
