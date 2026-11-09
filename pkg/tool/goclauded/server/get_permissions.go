package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetPermissions(
	x context.Context,
	r server.GetPermissionsRequestObject,
) (server.GetPermissionsResponseObject, error) {
	days := constant.PermissionWindowDays

	if r.Params.Days != nil {
		days = *r.Params.Days
	}

	report, e := s.transcript.Permissions(days)

	if e != nil {
		return server.GetPermissions500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	rows := []server.PermissionRow{}

	for _, w := range report.Rows {
		commands := []string{}
		commands = append(commands, w.Commands...)
		rows = append(
			rows,
			server.PermissionRow{
				Blocker:  w.Blocker,
				Prompts:  w.Prompts,
				Sessions: w.Sessions,
				Reason:   w.Reason,
				Last:     w.Last,
				Commands: commands,
			},
		)
	}

	return server.GetPermissions200JSONResponse{
		Days:     report.Days,
		Sessions: report.Sessions,
		Prompts:  report.Prompts,
		Declined: report.Declined,
		Cleared:  report.Cleared,
		Rows:     rows,
	}, nil
}
