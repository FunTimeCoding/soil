package coverage

import (
	"github.com/funtimecoding/soil/pkg/constant"
	crap "github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func TestPackages(
	root string,
	patterns ...string,
) []string {
	r := run.New()
	r.Directory = root

	return slice.StripEmpty(
		split.NewLine(
			strings.TrimSpace(
				r.Start(
					append(
						[]string{
							constant.Go,
							crap.ListCommand,
							crap.ListFormat,
							crap.TestPackageTemplate,
						},
						patterns...,
					)...,
				),
			),
		),
	)
}
