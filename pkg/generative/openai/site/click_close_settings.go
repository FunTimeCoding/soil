package site

import "github.com/funtimecoding/soil/pkg/generative/openai/constant"

func (s *Site) clickCloseSettings() {
	s.protocol.ClickQuery(constant.CloseSettingsSelector)
}
