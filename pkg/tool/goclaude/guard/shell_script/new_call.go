package shell_script

import (
	"mvdan.cc/sh/v3/syntax"
	"path"
)

func newCall(
	statement *syntax.Stmt,
	x *syntax.CallExpr,
	in *Script,
	nested *Script,
) *Call {
	readable := true

	for p := in; p != nil; p = p.Parent {
		if !p.Readable {
			readable = false
		}
	}

	words, fromInput := wrappedCommand(x.Args)
	name := ""

	if len(words) > 0 {
		program, _ := Argument(words[0])
		name = path.Base(program)
	}

	return &Call{
		Statement: statement,
		Expr:      x,
		Script:    in,
		Nested:    nested,
		Name:      name,
		Words:     words,
		FromInput: fromInput,
		Readable:  readable,
		TopLevel:  in.Parent == nil,
	}
}
