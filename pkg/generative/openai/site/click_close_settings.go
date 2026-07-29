package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickCloseSettings() {
	s.protocol.ClickQuery(constant.OpenAICloseSettingsSelector)
}
