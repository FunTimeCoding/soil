package site

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (s *Site) Send(t string) {
	if false {
		console.Format("Focused: %+v\n", s.protocol.Focused())

		return
	}

	s.protocol.EnterText(constant.OpenAIPromptSelector, t)
}
