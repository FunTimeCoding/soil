package reader

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (s *Reader) Dump() {
	console.Line(s.Protocol.Outer(constant.AnthropicBodyElement))
}
