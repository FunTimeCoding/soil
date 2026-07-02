package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) TableSizes(
	c context.Context,
	r server.TableSizesRequestObject,
) (server.TableSizesResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.TableSizes400JSONResponse(*clientError(e)), nil
	}

	schema := "public"

	if r.Params.Schema != nil {
		schema = *r.Params.Schema
	}

	rows, e := s.service.TableSizes(c, instance, schema)

	if e != nil {
		return server.TableSizes500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.TableSizes200JSONResponse{Rows: toRows(rows)}, nil
}
