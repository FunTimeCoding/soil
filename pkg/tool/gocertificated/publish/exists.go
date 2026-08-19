package publish

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func (p *Publisher) exists(path string) (bool, error) {
	_, e := p.forge.File(p.project, p.branch, path)

	if errors.Is(e, constant.ErrorNotFound) {
		return false, nil
	}

	return e == nil, e
}
