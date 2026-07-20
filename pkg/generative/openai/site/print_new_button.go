package site

import (
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/generative/openai/constant"
)

func (s *Site) printNewButton() {
	s.protocol.PrintNode(constant.NewSelector, constant.UsefulAttributes)
	protocol.Print(
		s.protocol.Select(constant.NewSelector, 0),
		constant.UsefulAttributes,
	)
}
