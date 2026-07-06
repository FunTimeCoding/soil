package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListVirtualMachines(
	_ context.Context,
	_ server.ListVirtualMachinesRequestObject,
) (server.ListVirtualMachinesResponseObject, error) {
	machines, e := s.client.VirtualMachines()

	if e != nil {
		return server.ListVirtualMachines500JSONResponse(*s.captureDetail(e)), nil
	}

	result := make([]*server.VirtualMachine, 0, len(machines))

	for _, m := range machines {
		entry := &server.VirtualMachine{
			Identifier: m.Identifier,
			Name:       m.Name,
		}

		if m.Raw.Cluster.IsSet() && m.Raw.Cluster.Get() != nil {
			entry.Cluster = new(m.Raw.Cluster.Get().GetName())
		}

		if m.Raw.Site.IsSet() && m.Raw.Site.Get() != nil {
			entry.Site = new(m.Raw.Site.Get().GetName())
		}

		if len(m.Tags) > 0 {
			entry.Tags = &m.Tags
		}

		result = append(result, entry)
	}

	return server.ListVirtualMachines200JSONResponse(result), nil
}
