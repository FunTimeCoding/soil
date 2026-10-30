package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service/oversize"
	"math"
	"testing"
)

func oversizeReport(
	t *testing.T,
	collection string,
) *oversize.Report {
	t.Helper()

	return oversizeReportWithin(t, collection, 5, math.MaxInt)
}
