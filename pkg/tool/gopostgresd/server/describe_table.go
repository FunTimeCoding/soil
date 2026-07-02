package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) DescribeTable(
	c context.Context,
	r server.DescribeTableRequestObject,
) (server.DescribeTableResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.DescribeTable400JSONResponse(*clientError(e)), nil
	}

	schema := "public"

	if r.Params.Schema != nil {
		schema = *r.Params.Schema
	}

	rows, e := s.service.DescribeTable(
		c,
		instance,
		schema,
		r.Table,
	)

	if e != nil {
		return server.DescribeTable500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.DescribeTable200JSONResponse{Rows: toRows(rows)}, nil
}
