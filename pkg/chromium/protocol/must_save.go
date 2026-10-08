package protocol

import "github.com/funtimecoding/soil/pkg/errors"

func (p *Protocol) MustSave(
	locator string,
	filename string,
) {
	errors.PanicOnError(p.Save(locator, filename))
}
