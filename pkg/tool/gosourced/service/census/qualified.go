package census

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/types"
)

func qualified(n *types.Named) string {
	return join.Empty(
		n.Obj().Pkg().Path(),
		constant.MemberSeparator,
		n.Obj().Name(),
	)
}
