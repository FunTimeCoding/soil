package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListAddressRanges(
	x context.Context,
	_ server.ListAddressRangesRequestObject,
) (server.ListAddressRangesResponseObject, error) {
	ranges, e := s.service.InternetAddressRanges()

	if unavailable.Is(e) {
		return server.ListAddressRanges503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListAddressRanges500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListAddressRanges200JSONResponse(
			convert.AddressRanges(ranges),
		),
		nil
}
