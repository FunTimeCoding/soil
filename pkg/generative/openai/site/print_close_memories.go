package site

import (
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (s *Site) printCloseMemories() {
	s.protocol.PrintNode(
		constant.OpenAICloseMemoriesSelector,
		constant.OpenAIUsefulAttributes,
	)
	n := s.protocol.Select(constant.OpenAICloseMemoriesSelector, 2)
	console.Line("Close dialog index 2")
	protocol.Print(n, constant.OpenAIUsefulAttributes)
}
