package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickProfile() {
	s.session.ClickQuery(constant.OpenAIProfileSelector)
}
