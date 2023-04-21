package parser

import (
  "testing"
  "ast"
  "lexer"
)

func checkParserErrors(t *testing.T, p *Parser) {
  errors := p.Errors()
  if len(errors) == 0 {
    return
  }

  t.Errorf("parser han %d errors", len(errors))
  for _, msg := range errors {
    t.Errorf("parser error: %q", msg)
  }
  t.FailNow()
}

func TestIntStatements(t *testing.T) {
  input := `
int x = 5;
int y = 10;
int foobar = 838383;
`

  l := lexer.New(input)
  p := New(l)

  program := p.ParseProgram()
  checkParserErrors(t, p)
  if program == nil {
    t.Fatalf("ParseProgram() returned nil")
  }
  if len(program.Statements) != 3 {
    t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
      3, len(program.Statements))
  }

  tests := []struct {
    expectedIdentifier string
  }{
    {"x"},
    {"y"},
    {"foobar"},
  }

  for i, tt := range tests {
    stmt := program.Statements[i]
    if !testIntStatement(t, stmt, tt.expectedIdentifier) {
      return
    }
  }
}

func testIntStatement(t *testing.T, s ast.Statement, name string) bool {
  if s.TokenLiteral() != "int" {
    t.Errorf("s.TokenLiteral not 'int'. got=%q", s.TokenLiteral())
    return false
  }

  intStmt, ok := s.(*ast.IntStatement)
  if !ok {
    t.Errorf("s not *ast.IntStatement. got=%T", s)
    return false
  }

  if intStmt.Name.Value != name {
    t.Errorf("intStmt.Name.Value not '%s'. got=%s", name, intStmt.Name.Value)
    return false
  }

  if intStmt.Name.TokenLiteral() != name {
    t.Errorf("intStmt.Name.TokenLiteral() not '%s'. got=%s", name, intStmt.Name.TokenLiteral())
    return false
  }

  return true
}
