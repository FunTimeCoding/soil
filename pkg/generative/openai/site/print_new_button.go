package site

import (
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (s *Site) printNewButton() {
	s.protocol.PrintNode(
		constant.OpenAINewSelector,
		constant.OpenAIUsefulAttributes,
	)
	protocol.Print(
		s.protocol.Select(constant.OpenAINewSelector, 0),
		constant.OpenAIUsefulAttributes,
	)
}
