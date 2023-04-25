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
int x;
int y;
int z = 42;
int foobar = 838383;
`

  l := lexer.New(input)
  p := New(l)

  program := p.ParseProgram()
  checkParserErrors(t, p)
  if program == nil {
    t.Fatalf("ParseProgram() returned nil")
  }
  if len(program.Statements) != 4 {
    t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
      4, len(program.Statements))
  }

  tests := []struct {
    expectedIdentifier string
  }{
    {"x"},
    {"y"},
    {"z"},
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

func TestReturnStatements(t *testing.T) {
  input := `
return 0;
return 1;
return 993322;
`

  l := lexer.New(input)
  p := New(l)

  program := p.ParseProgram()
  checkParserErrors(t, p)

  if len(program.Statements) != 3 {
    t.Fatalf("program.Statements does not contain 3 statements. got=%d",
      len(program.Statements))
  }

  for  _, stmt := range program.Statements {
    returnStmt, ok := stmt.(*ast.ReturnStatement)
    if !ok {
      t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
      continue
    }
    if returnStmt.TokenLiteral() != "return" {
      t.Errorf("returnStmt.TokenLiteral not 'return'. got=%q",
        returnStmt.TokenLiteral())
    }
  }
}
