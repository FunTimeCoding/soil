package protocol

import "github.com/funtimecoding/soil/pkg/errors"

func (p *Protocol) MustEnterText(
	s string,
	text string,
) {
	errors.PanicOnError(p.EnterText(s, text))
}
