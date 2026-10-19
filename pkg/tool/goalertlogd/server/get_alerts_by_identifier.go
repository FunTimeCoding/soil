package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
)

func (s *Server) GetAlertsByIdentifier(
	_ context.Context,
	r server.GetAlertsByIdentifierRequestObject,
) (server.GetAlertsByIdentifierResponseObject, error) {
	var identifiers []uint

	for _, i := range r.Params.Identifier {
		if i > 0 && len(identifiers) < constant.PageMaximum {
			identifiers = append(identifiers, uint(i))
		}
	}

	records, e := s.store.ByIdentifiers(identifiers)

	if e != nil {
		return server.GetAlertsByIdentifier500JSONResponse(*s.answer(e)), nil
	}

	return server.GetAlertsByIdentifier200JSONResponse(toResponse(records)), nil
}
