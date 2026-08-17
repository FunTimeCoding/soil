package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/forward"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Forwards(v []*forward.Forward) []server.Forward {
	result := []server.Forward{}

	for _, e := range v {
		result = append(result, *Forward(e))
	}

	return result
}
