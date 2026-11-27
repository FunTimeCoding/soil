package protocol

import "github.com/funtimecoding/soil/pkg/errors"

func (p *Protocol) MustClickQuery(s string) {
	errors.PanicOnError(p.ClickQuery(s))
}
