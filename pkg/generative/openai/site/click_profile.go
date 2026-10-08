package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickProfile() {
	s.session.MustClickQuery(constant.OpenAIProfileSelector)
}
