package reporter

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func functionName(qualified string) string {
	if i := strings.LastIndex(qualified, constant.Slash); i >= 0 {
		qualified = qualified[i+1:]
	}

	if i := strings.Index(qualified, constant.Dot); i >= 0 {
		return qualified[i+1:]
	}

	return qualified
}
