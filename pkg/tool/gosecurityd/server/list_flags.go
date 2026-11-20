package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) ListFlags(
	x context.Context,
	r server.ListFlagsRequestObject,
) (server.ListFlagsResponseObject, error) {
	all := r.Params.IncludeMitigated != nil && *r.Params.IncludeMitigated
	result, e := s.flag.Flags(all)

	if e != nil {
		return server.ListFlags500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ListFlags200JSONResponse(convert.Flags(result)), nil
}
