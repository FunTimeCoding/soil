package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostSessionLabel(
	_ context.Context,
	r server.PostSessionLabelRequestObject,
) (server.PostSessionLabelResponseObject, error) {
	var from string

	if r.Body.From != nil {
		from = *r.Body.From
	}

	var value string

	if r.Body.Value != nil {
		value = *r.Body.Value
	}

	var change string
	var e error

	switch {
	case r.Body.Expect != nil:
		change, e = s.service.SwapLabel(
			r.Identifier,
			from,
			r.Body.Key,
			*r.Body.Expect,
			value,
		)
	case value == "":
		change, e = s.service.DeleteLabel(r.Identifier, from, "", r.Body.Key)
	default:
		change, e = s.service.SetLabel(
			r.Identifier,
			from,
			"",
			r.Body.Key,
			value,
		)
	}

	if e != nil {
		if conflict.Is(e) {
			return server.PostSessionLabel409JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		return server.PostSessionLabel500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostSessionLabel200JSONResponse(
		server.LabelResponse{Change: change},
	), nil
}
