package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/machine_change"
)

func (s *Server) UpdateVirtualMachine(
	_ context.Context,
	r server.UpdateVirtualMachineRequestObject,
) (server.UpdateVirtualMachineResponseObject, error) {
	c := machine_change.New()
	c.Name = r.Body.Name
	c.PrimaryAddress = r.Body.PrimaryAddress
	c.Platform = r.Body.Platform
	c.Tenant = r.Body.Tenant
	c.Status = r.Body.Status

	if r.Body.Cores != nil {
		c.Cores = new(float64(*r.Body.Cores))
	}

	if r.Body.Memory != nil {
		c.Memory = new(int32(*r.Body.Memory))
	}

	m, e := s.service.UpdateVirtualMachine(r.Name, c)

	if e != nil {
		return server.UpdateVirtualMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.UpdateVirtualMachine200JSONResponse(
		*convert.VirtualMachine(m),
	), nil
}
