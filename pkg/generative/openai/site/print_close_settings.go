package site

import "github.com/funtimecoding/soil/pkg/generative/openai/constant"

func (s *Site) printCloseSettings() {
	s.protocol.PrintNode(constant.CloseSettingsSelector, constant.UsefulAttributes)
}
