package protocol

import "github.com/funtimecoding/soil/pkg/errors"

func (p *Protocol) MustOuter(s string) string {
	result, e := p.Outer(s)
	errors.PanicOnError(e)

	return result
}
