package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickMemories() {
	n := s.session.Select(constant.OpenAIMemoriesSelector, 0)

	if n == nil {
		return
	}

	s.session.ClickSearch(n.FullXPath())
}
