package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
)

func (s *Server) DeleteSessionById(
	_ context.Context,
	r server.DeleteSessionByIdRequestObject,
) (server.DeleteSessionByIdResponseObject, error) {
	confirm := ""

	if r.Params.Confirm != nil {
		confirm = *r.Params.Confirm
	}

	result, e := s.service.DeleteSession(r.Identifier, confirm)

	if e != nil {
		if refusal.Is(e) {
			return server.DeleteSessionById409JSONResponse{
				Error: e.Error(),
			}, nil
		}

		return server.DeleteSessionById500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.DeleteSessionById200JSONResponse(*deleteReceipt(result)), nil
}
