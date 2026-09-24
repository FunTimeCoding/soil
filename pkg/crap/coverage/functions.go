package coverage

import (
	"github.com/funtimecoding/soil/pkg/constant"
	crap "github.com/funtimecoding/soil/pkg/crap/constant"
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
			constant.Go,
			crap.Tool,
			crap.Cover,
			key_value.Equals(crap.FunctionReport, profile),
		),
	)
}
