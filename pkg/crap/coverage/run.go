package coverage

import (
	"github.com/funtimecoding/soil/pkg/constant"
	crap "github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
)

func Run(
	root string,
	profile string,
	patterns ...string,
) string {
	if profile == "" {
		profile = filepath.Join(root, crap.ProfileFile)
	}

	r := run.New()
	r.Directory = root
	r.Start(
		append(
			[]string{
				constant.Go,
				crap.TestCommand,
				key_value.Equals(crap.CoverProfile, profile),
				key_value.Equals(crap.CoverPackage, join.Comma(patterns)),
			},
			patterns...,
		)...,
	)

	return profile
}
