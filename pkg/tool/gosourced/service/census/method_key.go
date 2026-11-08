package census

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/types"
)

func methodKey(
	n *types.Named,
	method string,
) string {
	return join.Empty(qualified(n), constant.MemberSeparator, method)
}
