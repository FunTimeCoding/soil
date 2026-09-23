package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickPersonalize() {
	s.session.ClickQuery(constant.OpenAIPersonalizeSelector)
}
