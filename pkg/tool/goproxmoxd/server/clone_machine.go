package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func (s *Server) CloneMachine(
	x context.Context,
	r server.CloneMachineRequestObject,
) (server.CloneMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.CloneMachine404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.CloneMachine400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CloneMachine500JSONResponse(*s.answer(x, e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	options := &proxmox.VirtualMachineCloneOptions{Name: r.Body.Name}

	if r.Body.NewIdentifier != nil {
		options.NewID = int(*r.Body.NewIdentifier)
	}

	if r.Body.Full != nil && *r.Body.Full {
		options.Full = true
	}

	if r.Body.Storage != nil {
		options.Storage = *r.Body.Storage
	}

	if r.Body.Snapshot != nil {
		options.SnapName = *r.Body.Snapshot
	}

	newIdentifier, e := s.service.CloneMachine(
		instance,
		int(r.Identifier),
		node,
		options,
	)

	switch {
	case validation.Is(e):
		return server.CloneMachine400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.CloneMachine404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.CloneMachine403JSONResponse(*clientError(e)), nil
	case conflict.Is(e):
		return server.CloneMachine409JSONResponse(*clientError(e)), nil
	case unavailable.Is(e):
		return server.CloneMachine503JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.CloneMachine500JSONResponse(*s.answer(x, e)), nil
	}

	status := "cloned"

	return server.CloneMachine200JSONResponse{
		Identifier: int64(newIdentifier),
		Status:     &status,
	}, nil
}
