package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/server"
)

func (s *Server) RequestPower(
	_ context.Context,
	r server.RequestPowerRequestObject,
) (server.RequestPowerResponseObject, error) {
	instance, e := s.resolveInstance(r.Body.Instance)

	if not_found.Is(e) {
		return server.RequestPower404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.RequestPower400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.RequestPower500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	requested, f := s.service.Power(
		instance,
		string(r.Body.Operation),
		r.Body.Force != nil && *r.Body.Force,
	)

	if validation.Is(f) {
		return server.RequestPower400JSONResponse(*clientError(f)), nil
	}

	if conflict.Is(f) {
		return server.RequestPower409JSONResponse(*clientError(f)), nil
	}

	if f != nil {
		return server.RequestPower500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.RequestPower200JSONResponse{
		Instance:  instance,
		Requested: requested,
	}, nil
}
