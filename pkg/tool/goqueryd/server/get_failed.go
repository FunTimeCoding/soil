package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) GetFailed(
	x context.Context,
	r server.GetFailedRequestObject,
) (server.GetFailedResponseObject, error) {
	var collection string

	if r.Params.Collection != nil {
		collection = *r.Params.Collection
	}

	failed, e := s.service.Failed(collection)

	if e != nil {
		return server.GetFailed500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	chunks := make([]server.FailedChunk, len(failed))

	for i, c := range failed {
		chunks[i] = server.FailedChunk{
			Collection: c.Collection,
			Path:       c.Path,
			Index:      c.Index,
			FirstLine:  c.FirstLine,
			LastLine:   c.LastLine,
			Bytes:      c.Bytes,
			Reason:     c.Reason,
		}
	}

	return server.GetFailed200JSONResponse{Chunks: chunks}, nil
}
