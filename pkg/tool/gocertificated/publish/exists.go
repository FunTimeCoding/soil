package publish

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (p *Publisher) exists(path string) (bool, error) {
	_, e := p.forge.File(p.project, p.branch, path)

	if not_found.Is(e) {
		return false, nil
	}

	return e == nil, e
}
