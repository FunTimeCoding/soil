package lint

import (
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/repository"
)

func resolver(
	p *repository.Repository,
	o *option.Lint,
) *pointer.Resolver {
	r := pointer.New()
	r.Roots = pointer.Roots(p.Files.Files())
	r.ImplicitBases = p.ImplicitBases
	r.Registries = o.Registries
	r.Exists = p.Exists
	r.SiblingExists = p.SiblingExists
	r.Ignored = git.IgnoreMatcher(p.Root)
	r.Stdlib = stdlibMatcher()
	r.Dependency = dependencyMatcher(p.Modules)
	r.PrefixExists = p.PrefixExists
	r.Literal = p.Literal
	r.Routes = p.Routes

	return r
}
