package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Optional(constant.DefaultSpaceEnvironment); s != "" {
		o = append(o, WithDefaultSpace(s))
	}

	if s := environment.Optional(constant.DefaultPageEnvironment); s != "" {
		o = append(o, WithDefaultPage(s))
	}

	if v := environment.Slice(constant.LabelEnvironment); len(v) > 0 {
		o = append(o, WithLabel(v))
	}

	return New(
		environment.Required(atlassian.HostEnvironment),
		environment.Required(atlassian.UserEnvironment),
		environment.Required(atlassian.TokenEnvironment),
		o...,
	)
}
