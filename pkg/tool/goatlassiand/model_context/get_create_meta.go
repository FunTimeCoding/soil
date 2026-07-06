package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) getCreateMeta(
	_ context.Context,
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

	var expand []string

	if v := r.GetString(constant.ExpandFields, ""); v != "" {
		for _, s := range strings.Split(v, ",") {
			expand = append(expand, strings.TrimSpace(s))
		}
	}

	p, h := s.jira.MetaProject(project)

	if h != nil {
		return s.captureFail(h, "project metadata not found")
	}

	t, i := s.jira.MetaIssueType(p, issueType)

	if i != nil {
		return s.captureFail(i, "issue type not found")
	}

	return response.SuccessAny(convert.JiraCreateMeta(t, expand))
}
