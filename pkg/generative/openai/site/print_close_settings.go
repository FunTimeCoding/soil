package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) printCloseSettings() {
	s.protocol.PrintNode(
		constant.OpenAICloseSettingsSelector,
		constant.OpenAIUsefulAttributes,
	)
}
