package site

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (s *Site) Send(t string) {
	if false {
		fmt.Printf("Focused: %+v\n", s.protocol.Focused())

		return
	}

	s.protocol.EnterText(constant.OpenAIPromptSelector, t)
}
