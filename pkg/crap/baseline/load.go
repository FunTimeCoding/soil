package baseline

import (
	"github.com/funtimecoding/soil/pkg/crap/report"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
)

func Load(path string) *Baseline {
	var r report.Report
	notation.MustDecodeBytes(system.ReadBytesUnsafe(path), &r, false)

	return New(&r)
}
