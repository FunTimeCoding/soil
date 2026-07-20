package site

import "github.com/funtimecoding/soil/pkg/generative/openai/constant"

func (s *Site) printMemories() {
	s.protocol.PrintNode(constant.MemoriesSelector, []string{"class"})
}
