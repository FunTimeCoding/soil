package site

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"time"
)

func (s *Site) clickNew() {
	n := s.session.Select(constant.OpenAINewSelector, 0)

	if n == nil {
		return
	}

	s.session.ClickSearch(n.FullXPath())
	time.Sleep(1 * time.Second)
}
