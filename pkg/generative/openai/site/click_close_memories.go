package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickCloseMemories() {
	n := s.protocol.Select(constant.OpenAICloseMemoriesSelector, 2)

	if n == nil {
		return
	}

	s.protocol.ClickSearch(n.FullXPath())
}
