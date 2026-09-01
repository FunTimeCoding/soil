package sound

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/sound/constant"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"runtime"
)

func Play(
	path string,
	volume float64,
	verbose bool,
) {
	switch p := runtime.GOOS; p {
	case systemConstant.Linux:
		if verbose {
			console.Line("Sound not implemented on Linux")
		}
	case systemConstant.Darwin:
		result, e := system.RunError(
			constant.Afplay,
			constant.VolumeArgument,
			fmt.Sprintf("%.2f", volume),
			path,
		)

		if e != nil && verbose {
			console.Format("Sound error: %s\n", e)
			console.Format("Output: %s\n", result)
		}
	default:
		unexpected.String(p)
	}
}
