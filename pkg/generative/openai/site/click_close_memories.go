package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickCloseMemories() {
	n, okay := s.session.MustFindNode(
		constant.OpenAICloseMemoriesSelector,
		constant.OpenAICloseMemoriesIndex,
	)

	if !okay {
		return
	}

	s.session.MustClickSearch(n.FullXPath())
}
