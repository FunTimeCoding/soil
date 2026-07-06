package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) Explain(
	c context.Context,
	r server.ExplainRequestObject,
) (server.ExplainResponseObject, error) {
	instance, e := s.resolveInstance(r.Body.Instance)

	if e != nil {
		return server.Explain400JSONResponse(*clientError(e)), nil
	}

	prefix := "EXPLAIN"

	if r.Body.Analyze != nil && *r.Body.Analyze {
		prefix = "EXPLAIN ANALYZE"
	}

	rows, e := s.service.Query(
		c,
		instance,
		fmt.Sprintf("%s %s", prefix, r.Body.Sql),
	)

	if e != nil {
		return server.Explain500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.Explain200JSONResponse{Rows: toRows(rows)}, nil
}
