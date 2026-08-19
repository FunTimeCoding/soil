package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
)

func Changes(v []*publish.Change) []server.PendingChange {
	result := make([]server.PendingChange, 0, len(v))

	for _, c := range v {
		result = append(
			result,
			server.PendingChange{Path: c.Path, Reason: c.Reason},
		)
	}

	return result
}
