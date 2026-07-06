package model_context

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getIssue(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(parameter.Key)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	i, g := s.jira.Issue(key)

	if g != nil {
		return s.captureFail(g, "issue not found")
	}

	includeCustom := r.GetBool(constant.CustomFields, false)
	includeComments := r.GetBool(constant.Comments, false)
	var t *jira.MetaIssueType

	if includeCustom {
		p, h := s.jira.MetaProject(i.Raw.Fields.Project.Key)

		if h != nil {
			return s.captureFail(h, "project metadata not found")
		}

		t = p.GetIssueTypeWithName(i.Type)
	}

	result := convert.JiraIssueCustomFields(
		i,
		t,
		includeComments,
	)
	value := i.CustomValue(constant.ChecklistField)

	if value != "" &&
		value != issue.NilValue &&
		value != issue.UnknownField &&
		value != issue.UnknownValue {
		result.Checklist = ParseChecklist(value)
	}

	return response.SuccessAny(result)
}
