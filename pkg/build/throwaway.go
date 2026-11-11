package build

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/build/constant"
	soil "github.com/funtimecoding/soil/pkg/constant"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func throwaway(
	directory string,
	module string,
	version string,
	name string,
) error {
	system.SaveFile(
		join.Absolute(directory, goModConstant.ModFile),
		fmt.Sprintf("module %s\n", constant.ThrowawayModule),
	)
	r := run.New()
	r.Panic = false
	r.Directory = directory
	r.Start(
		soil.Go,
		soil.Get,
		fmt.Sprintf(
			"%s/%s/%s@%s",
			module,
			systemConstant.CommandPath,
			name,
			version,
		),
	)

	return r.Error
}
