package go_mod

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func ListModules(
	path string,
	modules []string,
	verbose bool,
) []*response.Module {
	if len(modules) == 0 {
		return nil
	}

	r := run.New()
	r.Directory = path
	r.Verbose = verbose
	r.Start(
		append(
			[]string{
				constant.Go,
				constant.List,
				constant.ModuleArgument,
				constant.UpdateArgument,
				constant.NotationArgument,
			},
			modules...,
		)...,
	)

	return notation.MustDecodeStream[response.Module](r.OutputString)
}
