package store_tester

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func Halves() string {
	return join.Empty(
		strings.Repeat("alfa\n", 800),
		strings.Repeat("bravo\n", 800),
	)
}
