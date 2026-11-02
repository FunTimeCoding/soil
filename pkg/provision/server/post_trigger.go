package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/generated/server"
	"github.com/funtimecoding/soil/pkg/provision/types/host_event"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
)

func (s *Server) PostTrigger(
	_ context.Context,
	r server.PostTriggerRequestObject,
) (server.PostTriggerResponseObject, error) {
	request := trigger.Request{
		Update:     r.Body.Update,
		Parameters: map[string]any{},
	}

	if len(r.Body.Changes) > 0 {
		request.Parameters[constant.DownstreamChanges] = r.Body.Changes
	}

	if r.Body.Hosts != nil {
		hosts := []*host_event.Event{}

		for _, h := range *r.Body.Hosts {
			hosts = append(hosts, host_event.New(h.Host, string(h.Action)))
		}

		request.Parameters[constant.DownstreamHosts] = hosts
	}

	if e := s.runner.Trigger(request); e != nil {
		if conflict.Is(e) {
			return server.PostTrigger409JSONResponse{Error: e.Error()}, nil
		}

		if unavailable.Is(e) {
			return server.PostTrigger503JSONResponse{Error: e.Error()}, nil
		}

		return server.PostTrigger500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	return server.PostTrigger202Response{}, nil
}
