package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
	"time"
)

func TestPermissionsHoldReportUntilRefresh(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	x := service_tester.New(t)
	project := t.TempDir()
	permissionSettings(t, project, "settings.local.json", "Bash(ls:*)")
	s := session.New("session-a")
	s.WorkDirectory = project
	x.Client.AddSession(s)
	x.Client.AddToolCall(
		"session-a",
		permissionCall(
			"Bash",
			"sort x",
			"user_temporary",
			"bashMiss:no-rule-match",
			time.Hour,
		),
	)
	setPermissionNow(x)
	r, e := x.Transcript.Permissions(7)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, r.Prompts)
	x.Client.AddToolCall(
		"session-a",
		permissionCall(
			"Bash",
			"sort y",
			"user_temporary",
			"bashMiss:no-rule-match",
			2*time.Hour,
		),
	)
	r, e = x.Transcript.Permissions(7)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, r.Prompts)
	r, e = x.Transcript.Permissions(8)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, r.Prompts)
	assert.Count(t, 0, x.Transcript.RefreshReports())
	r, e = x.Transcript.Permissions(7)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, r.Prompts)
}

func TestCoverageHoldsReportUntilRefresh(t *testing.T) {
	t.Setenv(constant.CoverageRootEnvironment, t.TempDir())
	x := service_tester.New(t)
	x.Client.AddSession(session.New("session-a"))
	now := time.Now().Format(time.RFC3339Nano)
	x.Client.AddToolCall("session-a", coverageCall("mcp__alfa__status", now))
	assert.Integer(t, 1, x.Transcript.Coverage()[0].CallsTotal)
	x.Client.AddToolCall("session-a", coverageCall("mcp__alfa__status", now))
	assert.Integer(t, 1, x.Transcript.Coverage()[0].CallsTotal)
	assert.Count(t, 0, x.Transcript.RefreshReports())
	assert.Integer(t, 2, x.Transcript.Coverage()[0].CallsTotal)
}
