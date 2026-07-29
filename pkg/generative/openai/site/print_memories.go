package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) printMemories() {
	s.protocol.PrintNode(constant.OpenAIMemoriesSelector, []string{"class"})
}
