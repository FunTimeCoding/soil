package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) printCloseSettings() {
	s.session.PrintNode(
		constant.OpenAICloseSettingsSelector,
		constant.OpenAIUsefulAttributes,
	)
}
