package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/generated/server"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
)

func (s *Server) PostTrigger(
	_ context.Context,
	r server.PostTriggerRequestObject,
) (server.PostTriggerResponseObject, error) {
	request := trigger.Request{Update: r.Body.Update}

	if len(r.Body.Changes) > 0 {
		request.Parameters = map[string]any{
			constant.DownstreamChanges: r.Body.Changes,
		}
	}

	if e := s.runner.Trigger(request); e != nil {
		if conflict.Is(e) {
			return server.PostTrigger409JSONResponse{Error: e.Error()}, nil
		}

		return server.PostTrigger500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	return server.PostTrigger202Response{}, nil
}
