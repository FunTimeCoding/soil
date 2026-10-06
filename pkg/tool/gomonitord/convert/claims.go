package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store/claim"
)

func Claims(v []*claim.Claim) []server.Claim {
	result := make([]server.Claim, 0, len(v))

	for _, c := range v {
		result = append(result, Claim(c))
	}

	return result
}
