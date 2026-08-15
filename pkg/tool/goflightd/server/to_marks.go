package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/mark"
)

func toMarks(v []mark.Mark) []server.MarkResponse {
	result := make([]server.MarkResponse, 0, len(v))

	for _, m := range v {
		result = append(result, *toMark(m))
	}

	return result
}
