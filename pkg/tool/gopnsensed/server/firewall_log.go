package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) FirewallLog(
	_ context.Context,
	r server.FirewallLogRequestObject,
) (server.FirewallLogResponseObject, error) {
	limit := constant.DefaultLogLimit

	if r.Params.Limit != nil {
		limit = *r.Params.Limit
	}

	result, e := s.opnsense.Log(limit)

	if e != nil {
		return server.FirewallLog500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.FirewallLog200JSONResponse(convert.LogEntries(result)), nil
}
