package relocator

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/types"
)

func QualifiedType(n *types.Named) string {
	return join.Dot([]string{n.Obj().Pkg().Name(), n.Obj().Name()})
}
