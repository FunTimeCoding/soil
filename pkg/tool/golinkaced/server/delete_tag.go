package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) DeleteTag(
	_ context.Context,
	r server.DeleteTagRequestObject,
) (server.DeleteTagResponseObject, error) {
	if e := s.client.DeleteTag(int(r.Identifier)); e != nil {
		return server.DeleteTag500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteTag200Response{}, nil
}
