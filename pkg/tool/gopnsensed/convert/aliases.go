package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/alias"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Aliases(v []*alias.Alias) []server.Alias {
	result := []server.Alias{}

	for _, e := range v {
		result = append(result, *Alias(e))
	}

	return result
}
