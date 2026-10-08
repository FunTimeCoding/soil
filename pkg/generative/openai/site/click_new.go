package site

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"time"
)

func (s *Site) clickNew() {
	n, okay := s.session.MustFindNode(
		constant.OpenAINewSelector,
		constant.OpenAINewIndex,
	)

	if !okay {
		return
	}

	s.session.MustClickSearch(n.FullXPath())
	time.Sleep(1 * time.Second)
}
