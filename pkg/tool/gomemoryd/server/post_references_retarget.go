package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) PostReferencesRetarget(
	_ context.Context,
	r server.PostReferencesRetargetRequestObject,
) (server.PostReferencesRetargetResponseObject, error) {
	reports, e := s.service.RetargetReferences(
		convert.MoveSet(r.Body),
		r.Body.DryRun,
	)

	if e != nil {
		return server.PostReferencesRetarget500JSONResponse(*s.answer(e)), nil
	}

	return server.PostReferencesRetarget200JSONResponse{
		Memories: convert.RetargetedMemories(reports),
	}, nil
}
