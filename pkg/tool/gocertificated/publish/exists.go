package publish

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (p *Publisher) exists(
	x context.Context,
	path string,
) (bool, error) {
	project, e := p.projectIdentifier(x)

	if e != nil {
		return false, e
	}

	_, e = p.forge.File(x, project, p.branch, path)

	if not_found.Is(e) {
		return false, nil
	}

	return e == nil, e
}
