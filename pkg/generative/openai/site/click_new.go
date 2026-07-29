package site

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"time"
)

func (s *Site) clickNew() {
	n := s.protocol.Select(constant.OpenAINewSelector, 0)

	if n == nil {
		return
	}

	s.protocol.ClickSearch(n.FullXPath())
	time.Sleep(1 * time.Second)
}
