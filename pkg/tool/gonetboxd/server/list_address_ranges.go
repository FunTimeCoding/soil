package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListAddressRanges(
	_ context.Context,
	_ server.ListAddressRangesRequestObject,
) (server.ListAddressRangesResponseObject, error) {
	ranges, e := s.service.InternetAddressRanges()

	if forbidden.Is(e) {
		return server.ListAddressRanges403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListAddressRanges500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListAddressRanges200JSONResponse(
			convert.AddressRanges(ranges),
		),
		nil
}
