package site

import "github.com/funtimecoding/soil/pkg/generative/openai/constant"

func (s *Site) clickCloseMemories() {
	n := s.protocol.Select(constant.CloseMemoriesSelector, 2)

	if n == nil {
		return
	}

	s.protocol.ClickSearch(n.FullXPath())
}
