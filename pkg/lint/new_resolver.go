package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/lint/repository"
)

func newResolver(
	p *repository.Repository,
	o *option.Lint,
) *resolver.Resolver {
	r := resolver.New()
	r.Roots = pointer.Roots(p.Files.Files())
	r.ImplicitBases = p.ImplicitBases
	r.Registries = o.Registries
	r.Exists = p.Exists
	r.SiblingExists = p.SiblingExists
	r.Ignored = p.Ignored
	r.Stdlib = stdlibMatcher()
	r.Dependency = dependencyMatcher(p.Modules)
	r.PrefixExists = p.PrefixExists
	r.Literal = p.Literal
	r.Routes = p.Routes
	r.Headings = p.Headings
	r.Declares = p.Declares

	return r
}
