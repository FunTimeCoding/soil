package unit

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/types/host_result"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/types/callback_report"
	"testing"
)

func TestCallbackReportReadsAnsibleOutput(t *testing.T) {
	var r callback_report.Report
	assert.FatalOnError(
		t,
		json.Unmarshal([]byte(constant.FixtureCallbackReport), &r),
	)
	assert.Any(
		t,
		[]*host_result.Result{
			host_result.New("alfa", "error", 3, "hello", "oops"),
			host_result.New(
				"bravo",
				"unreachable",
				0,
				"",
				"Task failed: Failed to connect to the host via ssh: ssh: connect to host 192.0.2.1 port 22: Operation timed out",
			),
		},
		r.Results(),
	)
}
