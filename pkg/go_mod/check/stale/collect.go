package stale

import (
	"github.com/funtimecoding/soil/pkg/go_mod"
	"github.com/funtimecoding/soil/pkg/go_mod/check/stale/option"
	"github.com/funtimecoding/soil/pkg/go_mod/module"
	"github.com/funtimecoding/soil/pkg/system"
	"time"
)

func collect(o *option.Stale) []*module.Module {
	if o.Path == "" {
		o.Path = system.WorkDirectory()
	}

	var paths []string

	for _, r := range go_mod.ReadDirectory(o.Path).Require {
		if r.Indirect && !o.Indirect {
			continue
		}

		paths = append(paths, r.Mod.Path)
	}

	now := time.Now()
	var result []*module.Module

	for _, r := range go_mod.ListModules(o.Path, paths, o.Verbose) {
		m := module.New(r, now, o.Months)
		m.Validate()
		result = append(result, m)
	}

	module.SortStalest(result)

	return result
}
