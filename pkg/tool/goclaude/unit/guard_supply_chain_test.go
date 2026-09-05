package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"testing"
)

func TestVerdictBlocksNpx(t *testing.T) {
	assert.String(
		t,
		"npx is blocked (supply-chain guard) - it downloads and executes npm packages on demand",
		guard.Verdict("darwin", "npx cowsay hi"),
	)
	assert.StringContains(
		t,
		"npx is blocked",
		guard.Verdict("darwin", "cd web && npx tsc --noEmit"),
	)
	assert.StringContains(
		t,
		"npx is blocked",
		guard.Verdict("darwin", "/usr/local/bin/npx serve ."),
	)
	assert.StringContains(
		t,
		"npx is blocked",
		guard.Verdict("linux", "npx create-react-app demo"),
	)
}

func TestVerdictBlocksPipInstall(t *testing.T) {
	assert.String(
		t,
		"pip install is blocked (supply-chain guard) - no python dependencies may be installed on this system",
		guard.Verdict("darwin", "pip install requests"),
	)
	assert.StringContains(
		t,
		"pip install is blocked",
		guard.Verdict("darwin", "pip3 install --upgrade requests"),
	)
	assert.StringContains(
		t,
		"pip install is blocked",
		guard.Verdict("darwin", "python -m pip install requests"),
	)
	assert.StringContains(
		t,
		"pip install is blocked",
		guard.Verdict("linux", "python3 -m pip install -U requests"),
	)
	assert.StringContains(
		t,
		"pip install is blocked",
		guard.Verdict("darwin", "pip --no-cache-dir install requests"),
	)
}

func TestVerdictAllowsPackageQueries(t *testing.T) {
	assert.String(t, "", guard.Verdict("darwin", "pip list"))
	assert.String(t, "", guard.Verdict("darwin", "pip3 show requests"))
	assert.String(t, "", guard.Verdict("darwin", "python -m venv .venv"))
	assert.String(t, "", guard.Verdict("darwin", "npm ls"))
	assert.String(
		t,
		"",
		guard.Verdict("darwin", `ssh host.example "npx cowsay hi"`),
	)
	assert.String(
		t,
		"",
		guard.Verdict("darwin", `ssh host.example "pip install requests"`),
	)
}
