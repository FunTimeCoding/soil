package coverage

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/crap/constant"
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
							soil.Go,
							constant.ListCommand,
							constant.ListFormat,
							constant.TestPackageTemplate,
						},
						patterns...,
					)...,
				),
			),
		),
	)
}
