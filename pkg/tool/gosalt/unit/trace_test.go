package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/salt/types/trace_report"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosalt/unit/command_tester"
	"testing"
)

func TestTraceReportsSummary(t *testing.T) {
	output, _, code := command_tester.Run(
		t,
		"trace",
		"web*",
		"--directory",
		fixture.Path(systemConstant.SaltPath, "state"),
	)
	assert.Integer(t, 0, code)
	var r trace_report.Report
	notation.MustDecode(output, &r, false)
	assert.String(t, "web*", r.Target)
	assert.Integer(t, 5, r.Summary.States)
	assert.Integer(t, 3, r.Summary.Files)
	assert.Integer(t, 4, r.Summary.OrphanedStates)
	assert.Integer(t, 2, r.Summary.OrphanedFiles)
}

func TestTraceUnknownTargetExits(t *testing.T) {
	_, failure, code := command_tester.Run(
		t,
		"trace",
		"nothing*",
		"--directory",
		fixture.Path(systemConstant.SaltPath, "state"),
	)
	assert.Integer(t, 1, code)
	assert.String(t, "top target not found: nothing*\n", failure)
}
