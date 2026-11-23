package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/salt/types/why_report"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosalt/unit/command_tester"
	"testing"
)

func TestWhyReportsChains(t *testing.T) {
	output, _, code := command_tester.Run(
		t,
		"why",
		"common.child",
		"--directory",
		fixture.Path(systemConstant.SaltPath, "state"),
	)
	assert.Integer(t, 0, code)
	var r why_report.Report
	notation.MustDecode(output, &r, false)
	assert.String(t, "common.child", r.State)
	assert.String(t, "reached", r.Status)
	assert.Count(t, 1, r.Chains)
	assert.Strings(t, []string{"common", "common.child"}, r.Chains[0].States)
}

func TestWhyHonoursEnvironmentFlag(t *testing.T) {
	output, _, code := command_tester.Run(
		t,
		"why",
		"stagingonly",
		"--directory",
		fixture.Path(systemConstant.SaltPath, "state"),
		"--environment",
		"base,prod,staging",
	)
	assert.Integer(t, 0, code)
	var r why_report.Report
	notation.MustDecode(output, &r, false)
	assert.String(t, "reached", r.Status)
	assert.True(t, r.Chains[0].Live)
}

func TestWhyUnknownStateExits(t *testing.T) {
	_, failure, code := command_tester.Run(
		t,
		"why",
		"nonexistent",
		"--directory",
		fixture.Path(systemConstant.SaltPath, "state"),
	)
	assert.Integer(t, 1, code)
	assert.String(t, "state not found: nonexistent\n", failure)
}
