package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclaude"
	"os"
	"path/filepath"
	"testing"
)

func TestLabelAllowanceFindsAllowLists(t *testing.T) {
	configuration := t.TempDir()
	project := t.TempDir()
	user := filepath.Join(configuration, "settings.json")
	local := filepath.Join(project, ".claude/settings.local.json")
	shared := filepath.Join(project, ".claude/settings.json")
	system.MakeDirectory(filepath.Dir(local))
	assert.FatalOnError(
		t,
		os.WriteFile(
			user,
			[]byte(`{"permissions":{"allow":["mcp__goclauded__label"]}}`),
			0600,
		),
	)
	assert.FatalOnError(
		t,
		os.WriteFile(
			local,
			[]byte(
				`{"permissions":{"allow":["mcp__plugin_soil_goclauded__*"]}}`,
			),
			0600,
		),
	)
	assert.FatalOnError(
		t,
		os.WriteFile(
			shared,
			[]byte(
				`{"permissions":{"allow":["mcp__goclauded__send","mcp__goclauded__labels"]}}`,
			),
			0600,
		),
	)
	assert.Strings(
		t,
		[]string{user, local},
		goclaude.LabelAllowance(configuration, project),
	)
}

func TestLabelAllowanceServerWide(t *testing.T) {
	configuration := t.TempDir()
	user := filepath.Join(configuration, "settings.local.json")
	assert.FatalOnError(
		t,
		os.WriteFile(
			user,
			[]byte(`{"permissions":{"allow":["mcp__goclauded"]}}`),
			0600,
		),
	)
	assert.Strings(
		t,
		[]string{user},
		goclaude.LabelAllowance(configuration, ""),
	)
}

func TestLabelAllowanceQuietWithoutSettings(t *testing.T) {
	assert.Strings(t, nil, goclaude.LabelAllowance(t.TempDir(), t.TempDir()))
}
