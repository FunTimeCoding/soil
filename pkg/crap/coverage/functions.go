package coverage

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func Functions(
	root string,
	profile string,
) map[string]float64 {
	r := run.New()
	r.Directory = root

	return Parse(
		r.Start(
			soil.Go,
			constant.Tool,
			constant.Cover,
			key_value.Equals(constant.FunctionReport, profile),
		),
	)
}
