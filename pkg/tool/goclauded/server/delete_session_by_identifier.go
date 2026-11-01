package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
)

func (s *Server) DeleteSessionByIdentifier(
	_ context.Context,
	r server.DeleteSessionByIdentifierRequestObject,
) (server.DeleteSessionByIdentifierResponseObject, error) {
	confirm := ""

	if r.Params.Confirm != nil {
		confirm = *r.Params.Confirm
	}

	result, e := s.service.DeleteSession(r.Identifier, confirm)

	if e != nil {
		if refusal.Is(e) {
			return server.DeleteSessionByIdentifier409JSONResponse{
				Error: e.Error(),
			}, nil
		}

		return server.DeleteSessionByIdentifier500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.DeleteSessionByIdentifier200JSONResponse(
		*deleteReceipt(result),
	), nil
}
