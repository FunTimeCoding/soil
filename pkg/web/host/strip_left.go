package host

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
)

func StripLeft(s string) string {
	if HasDot(s) {
		return join.Dot(split.Dot(s)[1:])
	}

	return s
}
