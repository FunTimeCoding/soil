package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) RunCron(
	_ context.Context,
	_ server.RunCronRequestObject,
) (server.RunCronResponseObject, error) {
	result, e := s.service.Cron()

	if forbidden.Is(e) {
		return server.RunCron403JSONResponse{Error: e.Error()}, nil
	}

	if conflict.Is(e) {
		return server.RunCron409JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.RunCron503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.RunCron500JSONResponse(*s.answer(e)), nil
	}

	return server.RunCron200JSONResponse(*convert.CronResult(result)), nil
}
