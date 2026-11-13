package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/repository"
)

func pointerChecker(
	repo *repository.Repository,
	o *option.Lint,
	r *output.Results,
) Checker {
	base := newResolver(repo, o)

	if o.Resolved == nil {
		return Pointers(base, r.AddUnchecked)
	}

	o.Resolved.Validate(repo.Root, pointerGlobal(repo, o, base))

	return cachedPointers(repo, o, r, base)
}
