package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) processReload(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ProcessReload,
) (*mcp.CallToolResult, error) {
	switch a.Scope {
	case constant.ProcfileScope:
		if e := s.supervisor.ReloadProcfile(); e != nil {
			return s.captureFail(e, "reload procfile failed")
		}
	case constant.EnvironmentScope:
		if e := s.supervisor.ReloadEnvironment(); e != nil {
			return s.captureFail(e, "reload environment failed")
		}
	default:
		return response.Fail(
			"scope must be %s or %s",
			constant.ProcfileScope,
			constant.EnvironmentScope,
		)
	}

	return response.SuccessAny(s.supervisor.Statuses())
}
