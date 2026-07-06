package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"time"
)

func (s *Server) GetWait(
	_ context.Context,
	r server.GetWaitRequestObject,
) (server.GetWaitResponseObject, error) {
	timeout := 30

	if r.Params.Timeout != nil {
		timeout = *r.Params.Timeout
	}

	pending, e := s.service.WaitMessage(
		r.Params.Callsign,
		time.Duration(timeout)*time.Second,
	)

	if e != nil {
		return server.GetWait500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	var result []server.Message

	for _, m := range pending {
		result = append(
			result,
			server.Message{
				From:      m.FromName,
				Body:      m.Body,
				Timestamp: m.CreatedAt.Format("2006-01-02T15:04:05Z"),
			},
		)
	}

	if result == nil {
		result = []server.Message{}
	}

	return server.GetWait200JSONResponse{Messages: result}, nil
}
