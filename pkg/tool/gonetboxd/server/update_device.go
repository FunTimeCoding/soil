package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/device_change"
)

func (s *Server) UpdateDevice(
	_ context.Context,
	r server.UpdateDeviceRequestObject,
) (server.UpdateDeviceResponseObject, error) {
	c := device_change.New()
	c.Name = r.Body.Name
	c.PrimaryAddress = r.Body.PrimaryAddress
	c.Location = r.Body.Location
	c.Platform = r.Body.Platform
	c.Tenant = r.Body.Tenant
	c.Status = r.Body.Status
	c.Serial = r.Body.Serial
	c.Description = r.Body.Description

	if r.Body.CustomFields != nil {
		c.CustomFields = *r.Body.CustomFields
	}

	d, e := s.service.UpdateDevice(r.Name, c)

	if not_found.Is(e) {
		return server.UpdateDevice404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.UpdateDevice500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.UpdateDevice200JSONResponse(*convert.Device(d)), nil
}
