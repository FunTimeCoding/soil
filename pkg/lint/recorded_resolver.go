package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/lint/types/pointer_dependency"
	"github.com/funtimecoding/soil/pkg/markup/heading"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path"
)

func recordedResolver(
	base *resolver.Resolver,
	d *pointer_dependency.Set,
	siblings []string,
) *resolver.Resolver {
	result := *base
	exact := func(p string) {
		clean := path.Clean(p)
		d.AddPath(clean)

		if outsideTrees(clean, siblings) {
			d.Outside = true
		}
	}
	under := func(directory string) {
		if directory == "" {
			d.AddPrefix("")

			return
		}

		clean := path.Clean(directory)
		exact(clean)
		d.AddPrefix(join.Empty(clean, stringsConstant.Slash))
	}
	result.Exists = func(p string) bool {
		exact(p)

		return base.Exists(p)
	}
	result.SiblingExists = func(p string) bool {
		exact(p)

		return base.SiblingExists(p)
	}
	result.PrefixExists = func(
		directory string,
		prefix string,
	) bool {
		under(directory)
		d.AddPrefix(join.Empty(directory, stringsConstant.Slash, prefix))

		return base.PrefixExists(directory, prefix)
	}
	result.Literal = func(
		directory string,
		needle string,
	) bool {
		under(directory)

		return base.Literal(directory, needle)
	}
	result.Routes = func(directory string) ([]string, bool) {
		exact(path.Join(directory, constant.RestSpecificationPath))

		return base.Routes(directory)
	}
	result.Headings = func(p string) ([]*heading.Heading, bool) {
		exact(p)

		return base.Headings(p)
	}
	result.Declares = func(
		directory string,
		name string,
	) bool {
		under(directory)

		return base.Declares(directory, name)
	}
	result.DeclaresType = func(
		directory string,
		name string,
	) bool {
		under(directory)

		return base.DeclaresType(directory, name)
	}
	result.Tests = func(name string) bool {
		d.Tests = true

		return base.Tests(name)
	}
	result.TestElsewhere = func(
		name string,
		directory string,
	) bool {
		d.Tests = true

		return base.TestElsewhere(name, directory)
	}

	return &result
}
