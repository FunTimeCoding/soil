package site

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (s *Site) Dump() {
	console.Line(s.protocol.Outer(constant.AnthropicBodyElement))
}
