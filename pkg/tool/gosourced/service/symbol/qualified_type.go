package symbol

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/types"
)

func qualifiedType(n *types.Named) string {
	return join.Dot([]string{n.Obj().Pkg().Name(), n.Obj().Name()})
}
