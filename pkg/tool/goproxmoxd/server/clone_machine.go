package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func (s *Server) CloneMachine(
	_ context.Context,
	r server.CloneMachineRequestObject,
) (server.CloneMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.CloneMachine400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.CloneMachine500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	options := &proxmox.VirtualMachineCloneOptions{
		Name: r.Body.Name,
	}

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

	newID, e := s.service.CloneMachine(
		c,
		int(r.Identifier),
		node,
		options,
	)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.CloneMachine404JSONResponse{
				Error: e.Error(),
			}, nil
		}

		return server.CloneMachine500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	status := "cloned"

	return server.CloneMachine200JSONResponse{
		Identifier: int64(newID),
		Status:     &status,
	}, nil
}
