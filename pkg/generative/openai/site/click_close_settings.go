package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickCloseSettings() {
	s.session.ClickQuery(constant.OpenAICloseSettingsSelector)
}
