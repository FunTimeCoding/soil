package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/server"
)

func (s *Server) DeleteEntry(
	_ context.Context,
	r server.DeleteEntryRequestObject,
) (server.DeleteEntryResponseObject, error) {
	if e := s.store.Delete(uint(r.Id)); e != nil {
		return server.DeleteEntry500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.DeleteEntry204Response{}, nil
}
