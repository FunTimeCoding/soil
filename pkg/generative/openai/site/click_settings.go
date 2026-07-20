package site

import "github.com/funtimecoding/soil/pkg/generative/openai/constant"

func (s *Site) clickSettings() {
	s.protocol.ClickQuery(constant.SettingsSelector)
}
