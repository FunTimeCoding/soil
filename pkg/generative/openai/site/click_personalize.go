package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickPersonalize() {
	s.protocol.ClickQuery(constant.OpenAIPersonalizeSelector)
}
