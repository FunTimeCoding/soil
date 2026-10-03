package site

import (
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (s *Site) printCloseMemories() {
	s.session.PrintNode(
		constant.OpenAICloseMemoriesSelector,
		constant.OpenAIUsefulAttributes,
	)
	n := s.session.Select(
		constant.OpenAICloseMemoriesSelector,
		constant.OpenAICloseMemoriesIndex,
	)
	console.Format("Close dialog index %d\n", constant.OpenAICloseMemoriesIndex)
	protocol.Print(n, constant.OpenAIUsefulAttributes)
}
