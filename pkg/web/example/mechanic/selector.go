package mechanic

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func selector(mark string) string {
	return join.Empty(constant.Hash, mark)
}
