package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_filter"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) ListMergeRequests(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListMergeRequests,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	since, e := parseSince(a.UpdatedAfter, time.Now())

	if e != nil {
		return s.answerMessage(x, e, "")
	}

	instance, project, e := s.resolveProject(x, a.Project)

	if e != nil {
		return s.answerMessage(x, e, "")
	}

	f := merge_request_filter.New(project)

	if a.State != "" {
		f.State = a.State
	}

	if a.Limit != 0 {
		f.Limit = a.Limit
	}

	if a.OrderBy != "" {
		f.OrderBy = a.OrderBy
	}

	f.Offset = a.Offset
	f.UpdatedAfter = since
	f.Author = a.Author
	f.Reviewer = a.Reviewer
	f.Assignee = a.Assignee
	f.SourceBranch = a.SourceBranch
	f.TargetBranch = a.TargetBranch
	f.Search = a.Search
	f.Draft = a.Draft
	v, e := s.service.MergeRequestsOf(x, instance, f)

	if e != nil {
		return s.answerMessage(x, e, "")
	}

	return response.SuccessAny(v)
}
