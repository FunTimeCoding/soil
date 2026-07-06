package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	jira "github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Optional(jira.DefaultProjectKeyEnvironment); s != "" {
		o = append(o, WithDefaultProjectKey(s))
	}

	if s := environment.Optional(
		jira.DefaultProjectNameEnvironment,
	); s != "" {
		o = append(o, WithDefaultProjectName(s))
	}

	if s := environment.Optional(jira.DefaultIssueTypeEnvironment); s != "" {
		o = append(o, WithDefaultIssueType(s))
	}

	if v := environment.Slice(jira.ClosedStatusEnvironment); len(v) > 0 {
		o = append(o, WithClosedStatus(v))
	}

	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.TokenEnvironment),
		o...,
	)
}
