package measure

import (
	"github.com/funtimecoding/soil/pkg/measure/constant"
	"github.com/funtimecoding/soil/pkg/measure/result"
)

func printUnplaced(r *result.Result) {
	printPaths(constant.UnplacedHeader, r.Unplaced)
	printPaths(constant.SkippedHeader, r.Skipped)
}
