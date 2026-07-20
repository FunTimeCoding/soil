package site

import "github.com/funtimecoding/soil/pkg/generative/openai/constant"

func (s *Site) clickMemories() {
	n := s.protocol.Select(constant.MemoriesSelector, 0)

	if n == nil {
		return
	}

	s.protocol.ClickSearch(n.FullXPath())
}
