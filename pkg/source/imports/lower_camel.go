package imports

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func lowerCamel(name string) string {
	first, rest, _ := strings.Cut(name, "_")

	return join.Empty(first, upperCamel(rest))
}
