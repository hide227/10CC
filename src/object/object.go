package object

import "fmt"

type ObjectType string

const (
  INTEGER_OBJ = "INTEGER"
  NULL_OBJ    = "NULL"
)

type Object interface {
  Type()    ObjectType
  Inspect() string
}

type Integer struct {
  Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type NULL struct {}

func (n *NULL) Type() ObjectType { return NULL_OBJ }
func (n *NULL) Inspect() string  { return "NULL" }
