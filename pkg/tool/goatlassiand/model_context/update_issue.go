package model_context

import (
	"context"
	"encoding/json"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) updateIssue(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(generative.ParameterKey)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	var labels []string

	if raw := r.GetString(constant.Labels, ""); raw != "" {
		if g := json.Unmarshal([]byte(raw), &labels); g != nil {
			return response.Fail(
				"labels must be a JSON array of strings: %v",
				g,
			)
		}
	}

	var fields map[string]any

	if raw := r.GetString(constant.AdditionalFields, ""); raw != "" {
		if g := json.Unmarshal([]byte(raw), &fields); g != nil {
			return response.Fail(
				"additional_fields must be a JSON object: %v",
				g,
			)
		}
	}

	result, g := s.service.UpdateIssue(
		key,
		r.GetString(constant.Summary, ""),
		r.GetString("description", ""),
		r.GetString(constant.Assignee, ""),
		r.GetString(constant.Reporter, ""),
		labels,
		fields,
	)

	if g != nil {
		return s.failOrCapture(g, "issue not updated")
	}

	return response.SuccessAny(
		convert.JiraIssueDiff(
			result.Before,
			result.After,
			r.GetBool(constant.NoDiff, false),
			result.CustomFieldNames,
		),
	)
}
