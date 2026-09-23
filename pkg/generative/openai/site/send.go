package site

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (s *Site) Send(t string) {
	if false {
		console.Format("Focused: %+v\n", s.session.Focused())

		return
	}

	s.session.EnterText(constant.OpenAIPromptSelector, t)
}
