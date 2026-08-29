package opnsense

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"slices"
)

func formatValidation(v map[string]string) string {
	var result []string

	for k, m := range v {
		result = append(result, key_value.ColonSpace(k, m))
	}

	slices.Sort(result)

	return join.CommaSpace(result)
}
