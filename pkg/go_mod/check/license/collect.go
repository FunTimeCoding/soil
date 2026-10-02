package license

import (
	"github.com/funtimecoding/soil/pkg/go_mod"
	"github.com/funtimecoding/soil/pkg/go_mod/check/license/option"
	"github.com/funtimecoding/soil/pkg/go_mod/dependency"
	"github.com/funtimecoding/soil/pkg/system"
)

func collect(o *option.License) []*dependency.Dependency {
	if o.Path == "" {
		o.Path = system.WorkDirectory()
	}

	result := go_mod.ListDependencies(o.Path, o.Verbose)

	for _, r := range result {
		r.Validate()
	}

	return result
}
