package object

import (
	"bytes"
	"strings"

	"github.com/exfly/newlang/ast"
)

// Function is the function type that holds the function's formal parameters,
// body and an environment to support closures.
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Bool() bool {
	return false
}

func (f *Function) String() string {
	return f.Inspect()
}

// Type returns the type of the object
func (f *Function) Type() Type { return FUNCTION }

// Inspect returns a stringified version of the object for debugging
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

// Return is the return type and used to hold the value of another object.
// This is used for `return` statements and this object is tracked through
// the evalulator and when encountered stops evaluation of the program,
// or body of a function.
type Return struct {
	Value Object
}

func (rv *Return) Bool() bool {
	return true
}

func (rv *Return) String() string {
	return rv.Inspect()
}

// Type returns the type of the object
func (rv *Return) Type() Type { return RETURN }

// Inspect returns a stringified version of the object for debugging
func (rv *Return) Inspect() string { return rv.Value.Inspect() }
