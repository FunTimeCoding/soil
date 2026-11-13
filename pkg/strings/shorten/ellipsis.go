package shorten

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Ellipsis(
	value string,
	count int,
) string {
	if len([]rune(value)) <= count {
		return value
	}

	return join.Empty(Prefix(value, count), constant.Ellipsis)
}
