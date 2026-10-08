package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickMemories() {
	n, okay := s.session.MustFindNode(constant.OpenAIMemoriesSelector, 0)

	if !okay {
		return
	}

	s.session.MustClickSearch(n.FullXPath())
}
