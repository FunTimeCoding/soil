package model_context

import (
	"context"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) resolveAssignee(
	c context.Context,
	name string,
) (*jira.User, *mcp.CallToolResult, error) {
	if name == constant.Unassign {
		return nil, nil, nil
	}

	users, _, e := s.jira.Nested().User.FindWithContext(c, name)

	if e != nil {
		r, f := response.Fail("assignee search failed: %v", e)

		return nil, r, f
	}

	if len(users) == 0 {
		r, f := response.Fail(
			"no user found matching: %s",
			name,
		)

		return nil, r, f
	}

	if len(users) == 1 {
		return &users[0], nil, nil
	}

	for i := range users {
		if users[i].DisplayName == name {
			return &users[i], nil, nil
		}
	}

	var names []string

	for _, u := range users {
		names = append(
			names,
			fmt.Sprintf("%s (%s)", u.DisplayName, u.AccountID),
		)
	}

	r, f := response.Fail(
		"multiple users match '%s': %v - be more specific or use an account ID",
		name,
		names,
	)

	return nil, r, f
}
