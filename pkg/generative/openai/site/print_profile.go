package site

import "github.com/funtimecoding/soil/pkg/generative/openai/constant"

func (s *Site) printProfile() {
	s.protocol.PrintNode(constant.ProfileSelector, constant.UsefulAttributes)
}
