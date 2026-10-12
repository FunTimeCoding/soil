package coverage

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
)

func Run(
	root string,
	profile string,
	patterns ...string,
) string {
	if profile == "" {
		directory := filepath.Join(root, systemConstant.Temporary)
		system.MakeDirectory(directory)
		profile = filepath.Join(directory, constant.ProfileFile)
	}

	r := run.New()
	r.Directory = root
	r.Start(
		append(
			[]string{
				soil.Go,
				constant.TestCommand,
				key_value.Equals(constant.CoverProfile, profile),
				key_value.Equals(constant.CoverPackage, join.Comma(patterns)),
			},
			patterns...,
		)...,
	)

	return profile
}
