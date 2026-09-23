package site

import (
	"github.com/funtimecoding/soil/pkg/chromium/session"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site/reader"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func New() *Site {
	s := session.New(constant.AnthropicUsageFragment)

	return &Site{Reader: reader.New(s.Protocol), session: s}
}
