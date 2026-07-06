package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) triggerRun(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	target := r.GetString(constant.Target, "")
	update := r.GetBool(constant.Update, false)
	synchronous := r.GetBool(constant.Synchronous, false)
	request := runner.TriggerRequest{Update: update}

	if target != "" {
		request.Parameters = map[string]any{
			constant.Target: target,
		}
	}

	if synchronous {
		request.Response = make(chan *runner.TriggerResult, 1)
	}

	e := s.runner.Trigger(request)

	if e != nil {
		return response.Fail(e.Error())
	}

	if synchronous {
		result := <-request.Response

		if result.Error != nil {
			return response.Fail(result.Error.Error())
		}

		return response.SuccessAny(result.Value)
	}

	return response.Success("run queued")
}
