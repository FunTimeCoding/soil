package site

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site/page"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (s *Site) ReadUsage() *page.Usage {
	return page.Parse(s.protocol.Outer(constant.AnthropicBodyElement))
}
