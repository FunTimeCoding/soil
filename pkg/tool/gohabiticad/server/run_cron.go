package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) RunCron(
	_ context.Context,
	_ server.RunCronRequestObject,
) (server.RunCronResponseObject, error) {
	result, e := s.habitica.Cron()

	if e != nil {
		return server.RunCron500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.RunCron200JSONResponse(
		*convert.CronResult(result),
	), nil
}
