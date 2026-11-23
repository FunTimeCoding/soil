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
	p := r.Params
	selected, e := convert.Selection(p.Show, p.Label, p.Search, p.Score, 0)

	if e != nil {
		return server.ListUpgrades400JSONResponse{Error: e.Error()}, nil
	}

	release := ""

	if p.Release != nil {
		release = *p.Release
	}

	result, f := s.system.Upgrades(
		selected,
		release,
		p.Kernel != nil && *p.Kernel,
	)

	if validation.Is(f) {
		return server.ListUpgrades400JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.ListUpgrades500JSONResponse(*s.answer(x, f)), nil
	}

	return server.ListUpgrades200JSONResponse(*convert.UpgradeList(result)), nil
}
