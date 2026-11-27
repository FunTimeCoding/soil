package protocol

import "github.com/funtimecoding/soil/pkg/errors"

func (p *Protocol) MustBody() string {
	result, e := p.Body()
	errors.PanicOnError(e)

	return result
}
