package protocol

import "github.com/funtimecoding/soil/pkg/errors"

func (p *Protocol) MustWaitVisible(s string) {
	errors.PanicOnError(p.WaitVisible(s))
}
