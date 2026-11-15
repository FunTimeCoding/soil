package shell_script

import "mvdan.cc/sh/v3/syntax"

type Call struct {
	Statement *syntax.Stmt
	Expr      *syntax.CallExpr
	Script    *Script
	Nested    *Script
	Name      string
	Words     []*syntax.Word
	FromInput bool
	Readable  bool
	TopLevel  bool
}
