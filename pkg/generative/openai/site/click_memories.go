package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickMemories() {
	n := s.protocol.Select(constant.OpenAIMemoriesSelector, 0)

	if n == nil {
		return
	}

	s.protocol.ClickSearch(n.FullXPath())
}
