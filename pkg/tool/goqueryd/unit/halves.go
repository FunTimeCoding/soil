package unit

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func halves() string {
	return join.Empty(
		strings.Repeat("alfa\n", 800),
		strings.Repeat("bravo\n", 800),
	)
}
