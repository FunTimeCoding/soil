package model_context

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) linkIssues(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(parameter.Key)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	target, g := r.RequireString(constant.TargetKey)

	if g != nil {
		return response.Fail("target_key is required: %v", g)
	}

	linkType := r.GetString(constant.LinkType, "Relates")
	link := &jira.IssueLink{
		Type:         jira.IssueLinkType{Name: linkType},
		OutwardIssue: &jira.Issue{Key: key},
		InwardIssue:  &jira.Issue{Key: target},
	}
	resp, h := s.jira.Nested().Issue.AddLinkWithContext(c, link)

	if h != nil {
		if resp != nil && resp.Body != nil {
			return response.Fail(
				"link failed: %s",
				string(system.ReadAll(resp.Body)),
			)
		}

		return response.Fail("link failed: %v", h)
	}

	return response.Success(
		"%s %s %s",
		key,
		linkType,
		target,
	)
}
