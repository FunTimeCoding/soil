package oversize_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/oversize"
	"math"
	"testing"
)

func Report(
	t *testing.T,
	collection string,
) *oversize.Report {
	t.Helper()

	return ReportWithin(t, collection, 5, math.MaxInt)
}
