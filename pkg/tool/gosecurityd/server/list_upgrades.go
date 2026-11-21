package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) ListUpgrades(
	x context.Context,
	r server.ListUpgradesRequestObject,
) (server.ListUpgradesResponseObject, error) {
	show := ""

	if r.Params.Show != nil {
		show = *r.Params.Show
	}

	result, e := s.system.Upgrades(show, convert.Minimum(r.Params.Score))

	if validation.Is(e) {
		return server.ListUpgrades400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListUpgrades500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ListUpgrades200JSONResponse(*convert.UpgradeList(result)), nil
}
