package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) clickSettings() {
	s.session.ClickQuery(constant.OpenAISettingsSelector)
}
