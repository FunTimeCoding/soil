package site

import "github.com/funtimecoding/soil/pkg/generative/openai/constant"

func (s *Site) clickProfile() {
	s.protocol.ClickQuery(constant.ProfileSelector)
}
