package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/server"
)

func (s *Server) GetStatus(
	_ context.Context,
	r server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.GetStatus400JSONResponse(*clientError(e)), nil
	}

	v, f := s.service.Status(instance)

	if f != nil {
		return server.GetStatus500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.GetStatus200JSONResponse{
		Instance:      v.Instance,
		Host:          v.Host,
		Provisioning:  v.Provisioning,
		ScreenEnabled: v.ScreenEnabled,
		ScreenTimeout: v.ScreenTimeout,
		Listener:      v.Listener,
		Redirection:   v.Redirection,
		ScreenService: v.ScreenService,
		Consent:       v.Consent,
		Power:         v.Power,
		Available:     v.Available,
		Operations:    toOperations(v.Operations),
	}, nil
}
