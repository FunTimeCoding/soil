package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/permission_tester"
	"testing"
	"time"
)

func TestClientReadsPermissionReport(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	s := base.New(t)
	project := t.TempDir()
	permission_tester.New(t).Settings(
		project,
		"settings.local.json",
		"Bash(ls:*)",
	)
	e := session.New("session-a")
	e.WorkDirectory = project
	s.Peers.Client.AddSession(e)
	c := tool_call.New(
		"Bash",
		"toolu_a",
		time.Now().Add(-time.Hour).Format(time.RFC3339Nano),
	)
	c.Detail = "sort x"
	c.PermissionDecision = "accept"
	c.PermissionSource = "user_temporary"
	c.PermissionReason = "bashMiss:no-rule-match"
	s.Peers.Client.AddToolCall("session-a", c)
	r, f := s.Client().GetPermissions(
		t.Context(),
		&client.GetPermissionsParams{},
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, 7, r.Days)
	assert.Integer(t, 1, r.Prompts)
	assert.Count(t, 1, r.Rows)
	assert.String(t, "sort", r.Rows[0].Blocker)
	assert.Strings(t, []string{"sort x"}, r.Rows[0].Commands)
}
