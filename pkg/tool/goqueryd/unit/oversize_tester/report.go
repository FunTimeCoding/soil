package oversize_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/oversize"
	"math"
)

func (o *Tester) Report(collection string) *oversize.Report {
	o.t.Helper()

	return o.ReportWithin(collection, 5, math.MaxInt)
}
