package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) GetUpgrade(
	x context.Context,
	r server.GetUpgradeRequestObject,
) (server.GetUpgradeResponseObject, error) {
	result, e := s.system.Upgrade(r.Release, r.Source)

	if not_found.Is(e) {
		return server.GetUpgrade404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetUpgrade500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetUpgrade200JSONResponse(convert.UpgradeGroup(result)), nil
}
