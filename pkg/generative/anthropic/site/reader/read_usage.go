package reader

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site/page"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (s *Reader) ReadUsage() *page.Usage {
	return page.Parse(s.Protocol.Outer(constant.AnthropicBodyElement))
}
