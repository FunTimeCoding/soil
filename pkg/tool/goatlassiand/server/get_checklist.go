package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetChecklist(
	_ context.Context,
	r server.GetChecklistRequestObject,
) (server.GetChecklistResponseObject, error) {
	items, e := s.service.ReadChecklist(r.Key)

	if e != nil {
		return server.GetChecklist500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.GetChecklist200JSONResponse(convertChecklist(items)), nil
}
