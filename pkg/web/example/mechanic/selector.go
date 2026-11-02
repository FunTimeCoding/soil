package mechanic

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func selector(mark string) string {
	return join.Empty(stringsConstant.Hash, mark)
}
