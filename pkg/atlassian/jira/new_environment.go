package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Optional(constant.JiraDefaultProjectKeyEnvironment); s != "" {
		o = append(o, WithDefaultProjectKey(s))
	}

	if s := environment.Optional(
		constant.JiraDefaultProjectNameEnvironment,
	); s != "" {
		o = append(o, WithDefaultProjectName(s))
	}

	if s := environment.Optional(constant.JiraDefaultIssueTypeEnvironment); s != "" {
		o = append(o, WithDefaultIssueType(s))
	}

	if v := environment.Slice(constant.JiraClosedStatusEnvironment); len(v) > 0 {
		o = append(o, WithClosedStatus(v))
	}

	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.TokenEnvironment),
		o...,
	)
}
