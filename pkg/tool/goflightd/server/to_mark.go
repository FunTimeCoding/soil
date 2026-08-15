package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/mark"
)

func toMark(v mark.Mark) *server.MarkResponse {
	result := &server.MarkResponse{
		Identifier: int64(v.Identifier),
		Time:       v.Time.Format(constant.DateFormat),
		Label:      v.Label,
	}

	if v.Note != "" {
		result.Note = new(v.Note)
	}

	return result
}
