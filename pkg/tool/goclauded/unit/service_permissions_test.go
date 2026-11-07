package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestPermissionsRankTheBlockersOfTheLastWeek(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	x := service_tester.New(t)
	project := t.TempDir()
	permissionSettings(
		t,
		project,
		"settings.local.json",
		"Bash(ls:*)",
		"Bash(cat:*)",
	)
	day := 24 * time.Hour

	for _, name := range []string{"session-a", "session-b", "session-c"} {
		s := session.New(name)
		s.WorkDirectory = project
		x.Client.AddSession(s)
	}

	c := x.Client
	c.AddToolCall(
		"session-a",
		permissionCall(
			"Bash",
			"cmp a b",
			"user_temporary",
			"bashMiss:no-rule-match",
			time.Hour,
		),
	)
	c.AddToolCall(
		"session-a",
		permissionCall(
			"Bash",
			"for f in x; do cmp $f y; done",
			"user_temporary",
			"bashMiss:too-complex",
			2*time.Hour,
		),
	)
	c.AddToolCall(
		"session-a",
		permissionCall("Bash", "ls doc", "config", "rule", time.Hour),
	)
	c.AddToolCall(
		"session-a",
		permissionCall(
			"Bash",
			"ls doc && cat x",
			"user_temporary",
			"bashMiss:too-complex",
			time.Hour,
		),
	)
	c.AddToolCall(
		"session-a",
		permissionCall("WebSearch", "", "user_temporary", "unknown", time.Hour),
	)
	c.AddToolCall(
		"session-a",
		permissionCall(
			"Bash",
			"rm doc/a.md",
			"user_reject",
			"bashMiss:no-rule-match",
			time.Hour,
		),
	)
	c.AddToolCall(
		"session-a",
		permissionCall(
			"AskUserQuestion",
			"",
			"user_temporary",
			"unknown",
			time.Hour,
		),
	)
	c.AddToolCall(
		"session-a",
		permissionCall(
			"Bash",
			"cmp old",
			"user_temporary",
			"bashMiss:no-rule-match",
			8*day,
		),
	)
	c.AddToolCall(
		"session-b",
		permissionCall(
			"Bash",
			"cmp c d",
			"user_temporary",
			"bashMiss:no-rule-match",
			3*time.Hour,
		),
	)
	c.AddToolCall(
		"session-c",
		permissionCall("Bash", "ls doc", "config", "rule", 9*day),
	)
	r, e := x.Transcript.Permissions(7, permissionNow())
	assert.FatalOnError(t, e)
	assert.Integer(t, 7, r.Days)
	assert.Integer(t, 2, r.Sessions)
	assert.Integer(t, 6, r.Prompts)
	assert.Integer(t, 1, r.Declined)
	assert.Integer(t, 1, r.Cleared)
	assert.Count(t, 3, r.Rows)
	assert.String(t, "cmp", r.Rows[0].Blocker)
	assert.Integer(t, 3, r.Rows[0].Prompts)
	assert.Integer(t, 2, r.Rows[0].Sessions)
	assert.String(t, "bashMiss:no-rule-match", r.Rows[0].Reason)
	assert.Time(t, permissionNow().Add(-time.Hour), r.Rows[0].Last)
	assert.Strings(
		t,
		[]string{"cmp a b", "for f in x; do cmp $f y; done", "cmp c d"},
		r.Rows[0].Commands,
	)
	assert.String(t, "WebSearch", r.Rows[1].Blocker)
	assert.Strings(t, nil, r.Rows[1].Commands)
	assert.String(t, "rm", r.Rows[2].Blocker)
}

func TestPermissionsMergeEverySettingsFile(t *testing.T) {
	home := t.TempDir()
	permissionSettings(t, home, "settings.json", "Bash(sort:*)")
	t.Setenv(soil.HomeEnvironment, home)
	x := service_tester.New(t)
	project := t.TempDir()
	permissionSettings(t, project, "settings.local.json", "Bash(jq:*)")
	permissionSettings(t, project, "settings.json", "Bash(awk:*)")
	s := session.New("session-a")
	s.WorkDirectory = project
	x.Client.AddSession(s)
	x.Client.AddToolCall(
		"session-a",
		permissionCall(
			"Bash",
			"jq . x | awk 1 | sort -u",
			"user_temporary",
			"bashMiss:too-complex",
			time.Hour,
		),
	)
	x.Client.AddToolCall(
		"session-a",
		permissionCall(
			"Bash",
			"rm -f tmp/work.md",
			"user_temporary",
			"bashMiss:no-rule-match",
			time.Hour,
		),
	)
	r, e := x.Transcript.Permissions(7, permissionNow())
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, r.Cleared)
	assert.Count(t, 0, r.Rows)
}

func TestPermissionsRowNamesItsNewestDistinctCommands(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	x := service_tester.New(t)
	project := t.TempDir()
	permissionSettings(t, project, "settings.local.json", "Bash(ls:*)")
	s := session.New("session-a")
	s.WorkDirectory = project
	x.Client.AddSession(s)

	for i, command := range []string{
		"ls doc | sort",
		"sort x",
		"sort x",
		"sort y",
		"sort z",
	} {
		reason := "bashMiss:too-complex"

		if i%2 == 1 {
			reason = "bashMiss:no-rule-match"
		}

		x.Client.AddToolCall(
			"session-a",
			permissionCall(
				"Bash",
				command,
				"user_temporary",
				reason,
				time.Duration(i+1)*time.Hour,
			),
		)
	}

	r, e := x.Transcript.Permissions(7, permissionNow())
	assert.FatalOnError(t, e)
	assert.Count(t, 1, r.Rows)
	assert.Integer(t, 5, r.Rows[0].Prompts)
	assert.String(t, "bashMiss:too-complex", r.Rows[0].Reason)
	assert.Strings(
		t,
		[]string{"ls doc | sort", "sort x", "sort y"},
		r.Rows[0].Commands,
	)
}

func TestPermissionsRefuseAnUnreadableSettingsFile(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	x := service_tester.New(t)
	project := t.TempDir()
	permissionSettings(t, project, "settings.local.json", "Bash(ls:*)")
	assert.FatalOnError(
		t,
		os.WriteFile(
			filepath.Join(project, ".claude", "settings.json"),
			[]byte("{"),
			0o644,
		),
	)
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
	_, e := x.Transcript.Permissions(7, permissionNow())
	assert.NotNil(t, e)
}

func TestPermissionsBreakAReasonTieAlphabetically(t *testing.T) {
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
			"bashMiss:too-complex",
			time.Hour,
		),
	)
	x.Client.AddToolCall(
		"session-a",
		permissionCall(
			"Bash",
			"sort y",
			"user_temporary",
			"bashMiss:no-rule-match",
			time.Hour,
		),
	)
	r, e := x.Transcript.Permissions(7, permissionNow())
	assert.FatalOnError(t, e)
	assert.String(t, "bashMiss:no-rule-match", r.Rows[0].Reason)
}
