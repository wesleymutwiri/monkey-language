package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
)

type String struct {
	Value string
}
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// type Environment struct {
// 	store map[string]Object
// }

type Error struct {
	Message string
}

type ReturnValue struct {
	Value Object
}

type Null struct{}

type Boolean struct {
	Value bool
}

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

func (n *Null) Type() ObjectType { return NULL_OBJ }

func (n *Null) Inspect() string { return "null" }

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

// func NewEnvironment() *Environment {
// 	s := make(map[string]Object)
// 	return &Environment{store: s}
// }

// func (e *Environment) Get(name string) (Object, bool) {
// 	obj, ok := e.store[name]
// 	return obj, ok
// }

// func (e *Environment) Set(name string, val Object) Object {
// 	e.store[name] = val
// 	return val
// }

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }
