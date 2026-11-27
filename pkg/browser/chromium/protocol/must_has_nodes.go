package protocol

import "github.com/funtimecoding/soil/pkg/errors"

func (p *Protocol) MustHasNodes(s string) bool {
	result, e := p.HasNodes(s)
	errors.PanicOnError(e)

	return result
}
