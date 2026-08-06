package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateAddressRange(
	_ context.Context,
	r server.CreateAddressRangeRequestObject,
) (server.CreateAddressRangeResponseObject, error) {
	status := ""

	if r.Body.Status != nil {
		status = *r.Body.Status
	}

	description := ""

	if r.Body.Description != nil {
		description = *r.Body.Description
	}

	result, e := s.client.CreateInternetAddressRange(
		r.Body.Start,
		r.Body.End,
		status,
		description,
	)

	if e != nil {
		return server.CreateAddressRange500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateAddressRange201JSONResponse(
		*convert.AddressRange(result),
	), nil
}
