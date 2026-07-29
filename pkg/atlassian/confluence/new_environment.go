package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Optional(constant.ConfluenceDefaultSpaceEnvironment); s != "" {
		o = append(o, WithDefaultSpace(s))
	}

	if s := environment.Optional(constant.ConfluenceDefaultPageEnvironment); s != "" {
		o = append(o, WithDefaultPage(s))
	}

	if v := environment.Slice(constant.ConfluenceLabelEnvironment); len(v) > 0 {
		o = append(o, WithLabel(v))
	}

	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.TokenEnvironment),
		o...,
	)
}
