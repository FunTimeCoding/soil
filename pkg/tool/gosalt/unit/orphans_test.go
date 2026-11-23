package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/salt/types/orphan_report"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosalt/unit/command_tester"
	"testing"
)

func TestOrphansReportsLiveEnvironments(t *testing.T) {
	output, _, code := command_tester.Run(
		t,
		"orphans",
		"--directory",
		fixture.Path(systemConstant.SaltPath, "state"),
	)
	assert.Integer(t, 0, code)
	var r orphan_report.Report
	notation.MustDecode(output, &r, false)
	assert.Strings(t, []string{"orphan", "stagingonly"}, r.Unreached)
	assert.Strings(
		t,
		[]string{"unused/draft.txt", "unused/notes.txt"},
		r.UnreachedFiles,
	)
	assert.Strings(
		t,
		[]string{"helpers.gamma", "service.alpha", "service.beta"},
		r.Possible,
	)
	assert.Strings(t, []string{"absent"}, r.Missing)
	assert.Strings(t, []string{"staging"}, r.DeadEnvironments)
	assert.Integer(t, 5, r.Summary.Reached)
	assert.Integer(t, 3, r.Summary.Possible)
	assert.Integer(t, 2, r.Summary.Unreached)
	assert.Integer(t, 1, r.Summary.Missing)
	assert.Integer(t, 2, r.Summary.UnreachedFiles)
}

func TestOrphansCountsPillarFiles(t *testing.T) {
	output, _, code := command_tester.Run(
		t,
		"orphans",
		"--directory",
		fixture.Path(systemConstant.SaltPath, "state"),
		"--pillar",
		fixture.Path(systemConstant.SaltPath, "pillar_root"),
	)
	assert.Integer(t, 0, code)
	var r orphan_report.Report
	notation.MustDecode(output, &r, false)
	assert.Strings(t, []string{"unused/draft.txt"}, r.UnreachedFiles)
}

func TestOrphansMergesDirectories(t *testing.T) {
	output, _, code := command_tester.Run(
		t,
		"orphans",
		"--directory",
		fixture.Path(systemConstant.SaltPath, "secret"),
		"--directory",
		fixture.Path(systemConstant.SaltPath, "state"),
	)
	assert.Integer(t, 0, code)
	var r orphan_report.Report
	notation.MustDecode(output, &r, false)
	assert.Count(t, 0, r.Missing)
	assert.Strings(t, []string{"stagingonly"}, r.Unreached)
}

func TestOrphansHonoursEnvironmentFlag(t *testing.T) {
	output, _, code := command_tester.Run(
		t,
		"orphans",
		"--directory",
		fixture.Path(systemConstant.SaltPath, "state"),
		"--environment",
		"base,prod,staging",
	)
	assert.Integer(t, 0, code)
	var r orphan_report.Report
	notation.MustDecode(output, &r, false)
	assert.Strings(t, []string{"orphan"}, r.Unreached)
	assert.Count(t, 0, r.DeadEnvironments)
}
