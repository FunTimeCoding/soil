package site

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Site) printMemories() {
	s.session.PrintNode(constant.OpenAIMemoriesSelector, []string{"class"})
}
