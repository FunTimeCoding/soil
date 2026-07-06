package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
)

func (s *Server) GetEntries(
	_ context.Context,
	r server.GetEntriesRequestObject,
) (server.GetEntriesResponseObject, error) {
	f := store.NewFilter()

	if r.Params.System != nil {
		f.System = *r.Params.System
	}

	if r.Params.Service != nil {
		f.Service = *r.Params.Service
	}

	if r.Params.User != nil {
		f.User = *r.Params.User
	}

	if r.Params.Since != nil {
		f.Since = *r.Params.Since
	}

	if r.Params.Until != nil {
		f.Until = *r.Params.Until
	}

	if r.Params.Limit != nil {
		f.Limit = *r.Params.Limit
	}

	entries, e := s.store.List(f)

	if e != nil {
		return server.GetEntries500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.GetEntries200JSONResponse(toResponse(entries)), nil
}
