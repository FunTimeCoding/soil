package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickCloseMemories() {
	n := s.session.Select(
		constant.OpenAICloseMemoriesSelector,
		constant.OpenAICloseMemoriesIndex,
	)

	if n == nil {
		return
	}

	s.session.ClickSearch(n.FullXPath())
}
