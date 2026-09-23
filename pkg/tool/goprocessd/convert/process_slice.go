package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/status"
)

func ProcessSlice(v []*status.Status) []server.Process {
	result := []server.Process{}

	for _, e := range v {
		result = append(result, *Process(e))
	}

	return result
}
