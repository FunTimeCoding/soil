package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) CreateMergeRequestNote(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateMergeRequestNote,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.MergeRequest == 0 {
		return response.Fail("merge_request is required")
	}

	if a.Body == "" {
		return response.Fail("body is required")
	}

	v, _, e := s.client.Notes.CreateMergeRequestNote(
		a.Project,
		a.MergeRequest,
		&gitlab.CreateMergeRequestNoteOptions{Body: &a.Body},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
