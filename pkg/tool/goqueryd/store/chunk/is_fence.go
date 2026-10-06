package chunk

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"strings"
)

func isFence(text string) bool {
	return strings.HasPrefix(text, constant.BacktickFence) ||
		strings.HasPrefix(text, constant.TildeFence)
}
