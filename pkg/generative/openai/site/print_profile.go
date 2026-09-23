package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) printProfile() {
	s.session.PrintNode(
		constant.OpenAIProfileSelector,
		constant.OpenAIUsefulAttributes,
	)
}
