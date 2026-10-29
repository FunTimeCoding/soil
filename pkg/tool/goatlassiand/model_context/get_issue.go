package model_context

import (
	"context"
	"github.com/andygrunwald/go-jira"
	atlassianConstant "github.com/funtimecoding/soil/pkg/atlassian/constant"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getIssue(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(generativeConstant.ParameterKey)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	i, g := s.service.Issue(key)

	if g != nil {
		return s.answer(g)
	}

	includeCustom := r.GetBool(constant.CustomFields, false)
	includeComments := r.GetBool(constant.Comments, false)
	var t *jira.MetaIssueType

	if includeCustom {
		var h error
		t, h = s.service.IssueType(i)

		if h != nil {
			return s.answerWith(h, "read project metadata")
		}
	}

	result := convert.JiraIssueCustomFields(i, t, includeComments)
	value := i.CustomValue(constant.ChecklistField)

	if value != "" &&
		value != atlassianConstant.JiraNilValue &&
		value != atlassianConstant.JiraUnknownField &&
		value != atlassianConstant.JiraUnknownValue {
		result.Checklist = checklist_item.Parse(value)
	}

	return response.SuccessAny(result)
}
