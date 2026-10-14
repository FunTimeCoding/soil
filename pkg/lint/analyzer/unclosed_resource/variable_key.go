package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/types"
)

func variableKey(o *types.Var) string {
	return join.Dot([]string{o.Pkg().Path(), o.Name()})
}
