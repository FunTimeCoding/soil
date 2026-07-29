package variable_naming

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/types"
)

type typedVariable struct {
	ident           *ast.Ident
	typ             types.Type
	precedence      int
	scopedNames     map[string]bool
	descendantNames map[string]bool
	kind            constant.VariableKind
	exempt          bool
}
