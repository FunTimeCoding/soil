package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) Reconcile(
	x context.Context,
	r server.ReconcileRequestObject,
) (server.ReconcileResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.Reconcile404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.Reconcile400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.Reconcile500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	p, e := s.service.Reconcile(
		x,
		instance,
		r.Params.Full != nil && *r.Params.Full,
		r.Params.Write != nil && *r.Params.Write,
	)

	if unavailable.Is(e) {
		return server.Reconcile503JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.Reconcile500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.Reconcile200JSONResponse(convertPlan(p)), nil
}
