package revision

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func packageKey(
	directory string,
	clause string,
) string {
	return join.Empty(directory, constant.KeySeparator, clause)
}
