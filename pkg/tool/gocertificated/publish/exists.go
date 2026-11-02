package publish

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (p *Publisher) exists(
	x context.Context,
	path string,
) (bool, error) {
	_, e := p.forge.File(x, p.project, p.branch, path)

	if not_found.Is(e) {
		return false, nil
	}

	return e == nil, e
}
