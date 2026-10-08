package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) Send(t string) {
	s.session.MustEnterText(constant.OpenAIPromptSelector, t)
}
