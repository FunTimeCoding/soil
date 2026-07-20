package site

import "github.com/funtimecoding/soil/pkg/generative/openai/constant"

func (s *Site) clickPersonalize() {
	s.protocol.ClickQuery(constant.PersonalizeSelector)
}
