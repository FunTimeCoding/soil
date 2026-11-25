package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	gitlabResponse "github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SubscribeMergeRequest(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.SubscribeMergeRequest,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.MergeRequest == 0 {
		return response.Fail("merge_request is required")
	}

	if a.Callsign == "" {
		return response.Fail("callsign is required")
	}

	d := goclaudedConstant.Delivery(a.Delivery)

	if d == "" {
		d = goclaudedConstant.DeliveryImmediate
	}

	if d != goclaudedConstant.DeliveryImmediate &&
		d != goclaudedConstant.DeliveryQueued {
		return response.Fail(
			"delivery must be immediate or queued: %s",
			a.Delivery,
		)
	}

	instance, project, e := s.resolveProject(x, a.Project)

	if e != nil {
		return s.answerMessage(x, e, "")
	}

	v, detail, existing, f := s.service.SubscribeMergeRequest(
		x,
		instance,
		project,
		a.MergeRequest,
		a.Callsign,
		a.Alias,
		d,
	)

	if f != nil {
		return s.answerMessage(x, f, "")
	}

	if existing {
		return response.Success("already subscribed")
	}

	return response.SuccessAny(gitlabResponse.NewWatched(v, detail))
}
