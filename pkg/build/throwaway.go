package build

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/build/constant"
	constant1 "github.com/funtimecoding/soil/pkg/constant"
	constant2 "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/system"
	constant3 "github.com/funtimecoding/soil/pkg/system/constant"
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
		join.Absolute(directory, constant2.ModFile),
		fmt.Sprintf("module %s\n", constant.ThrowawayModule),
	)
	r := run.New()
	r.Panic = false
	r.Directory = directory
	r.Start(
		constant1.Go,
		constant1.Get,
		fmt.Sprintf(
			"%s/%s/%s@%s",
			module,
			constant3.CommandPath,
			name,
			version,
		),
	)

	return r.Error
}
