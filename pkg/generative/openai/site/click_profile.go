package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickProfile() {
	s.protocol.ClickQuery(constant.OpenAIProfileSelector)
}
