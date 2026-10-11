package technitium

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/technitium/constant"
)

func listValue(v []string) string {
	if len(v) == 0 {
		return constant.ClearValue
	}

	return join.Comma(v)
}
