package unit

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"path/filepath"
	"testing"
	"time"
)

func TestCostWindowRunsOnServiceClock(t *testing.T) {
	s := base.New(t)
	s.Store.Advance(-30 * 24 * time.Hour)
	at := s.Store.Clock()().Add(-time.Hour).Format(time.RFC3339Nano)
	line := `{"uuid":"c1","type":"assistant","timestamp":"%s","requestId":"r1","message":{"id":"m1","role":"assistant","model":"claude-sonnet-4-5","content":[],"usage":{"input_tokens":10,"output_tokens":5}}}`
	system.WriteFile(
		filepath.Join(s.Harbor, "clocked.jsonl"),
		[]byte(join.Empty(fmt.Sprintf(line, at), "\n")),
		0644,
	)
	days := 1
	r, e := s.RESTClient(t).GetCostWithResponse(
		context.Background(),
		&client.GetCostParams{Days: &days},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.FatalNotNil(t, r.JSON200.Models)
	assert.Count(t, 1, *r.JSON200.Models)
	assert.String(t, "sonnet", (*r.JSON200.Models)[0].Model)
}

func TestPermissionWindowRunsOnServiceClock(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	s := base.New(t)
	s.Store.Advance(-30 * 24 * time.Hour)
	project := t.TempDir()
	permissionSettings(t, project, "settings.local.json", "Bash(ls:*)")
	e := session.New("session-a")
	e.WorkDirectory = project
	s.Set.Client.AddSession(e)
	c := tool_call.New(
		"Bash",
		"toolu_a",
		s.Store.Clock()().Add(-time.Hour).Format(time.RFC3339Nano),
	)
	c.Detail = "sort x"
	c.PermissionDecision = "accept"
	c.PermissionSource = "user_temporary"
	c.PermissionReason = "bashMiss:no-rule-match"
	s.Set.Client.AddToolCall("session-a", c)
	r, f := s.Client(t).GetPermissions(
		t.Context(),
		&client.GetPermissionsParams{},
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, r.Prompts)
}
