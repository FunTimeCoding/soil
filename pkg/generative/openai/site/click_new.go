package site

import (
	"github.com/funtimecoding/soil/pkg/generative/openai/constant"
	"time"
)

func (s *Site) clickNew() {
	n := s.protocol.Select(constant.NewSelector, 0)

	if n == nil {
		return
	}

	s.protocol.ClickSearch(n.FullXPath())
	time.Sleep(1 * time.Second)
}
