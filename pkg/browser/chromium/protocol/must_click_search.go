package protocol

import "github.com/funtimecoding/soil/pkg/errors"

func (p *Protocol) MustClickSearch(s string) {
	errors.PanicOnError(p.ClickSearch(s))
}
