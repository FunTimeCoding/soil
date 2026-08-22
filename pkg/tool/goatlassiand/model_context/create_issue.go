package model_context

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createIssue(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	project, f := r.RequireString(constant.Project)

	if f != nil {
		return response.Fail("project is required: %v", f)
	}

	issueType, g := r.RequireString(constant.IssueType)

	if g != nil {
		return response.Fail("issue_type is required: %v", g)
	}

	summary, h := r.RequireString(constant.Summary)

	if h != nil {
		return response.Fail("summary is required: %v", h)
	}

	var labels []string

	if raw := r.GetString(constant.Labels, ""); raw != "" {
		if i := json.Unmarshal([]byte(raw), &labels); i != nil {
			return response.Fail(
				"labels must be a JSON array of strings: %v",
				i,
			)
		}
	}

	var fields map[string]any

	if raw := r.GetString(constant.AdditionalFields, ""); raw != "" {
		if i := json.Unmarshal([]byte(raw), &fields); i != nil {
			return response.Fail(
				"additional_fields must be a JSON object: %v",
				i,
			)
		}
	}

	created, i := s.service.CreateIssue(
		project,
		issueType,
		summary,
		r.GetString("description", ""),
		r.GetString(constant.Assignee, ""),
		labels,
		fields,
	)

	if i != nil {
		return s.failOrCapture(i, "issue not created")
	}

	return response.SuccessAny(convert.JiraIssue(created))
}
