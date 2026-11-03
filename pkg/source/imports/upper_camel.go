package imports

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func upperCamel(name string) string {
	var result []string

	for _, s := range strings.Split(name, "_") {
		if s != "" {
			result = append(result, upperFirst(s))
		}
	}

	return join.Empty(result...)
}
