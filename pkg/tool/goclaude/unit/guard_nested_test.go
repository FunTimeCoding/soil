package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"testing"
)

func TestVerdictBlocksSupplyChainInsideAShellString(t *testing.T) {
	assert.String(
		t,
		"npx is blocked (supply-chain guard) - it downloads and executes npm packages on demand",
		guard.Verdict("darwin", "bash -c 'npx cowsay hi'"),
	)
	assert.String(
		t,
		"pip install is blocked (supply-chain guard) - no python dependencies may be installed on this system",
		guard.Verdict("linux", `sh -c "pip install $package"`),
	)
	assert.String(
		t,
		"npx is blocked (supply-chain guard) - it downloads and executes npm packages on demand",
		guard.Verdict("linux", "eval npx serve ."),
	)
}

func TestVerdictBlocksBsdFlagsInsideAShellString(t *testing.T) {
	assert.String(
		t,
		"sed on macOS is BSD sed and its flags (notably -i) differ from GNU sed - use gsed instead",
		guard.Verdict("darwin", "zsh -c 'sed -i s/a/b/ file.txt'"),
	)
	assert.String(
		t,
		"xargs on macOS is BSD xargs and rejects the GNU flags -a and -d - redirect the file on stdin instead: xargs command < file",
		guard.Verdict("darwin", "bash -lc 'xargs -a list.txt echo'"),
	)
}

func TestVerdictLeavesRemoteShellStrings(t *testing.T) {
	assert.String(
		t,
		"",
		guard.Verdict("darwin", `kubectl exec pod -- sh -c "npx serve ."`),
	)
	assert.String(t, "", guard.Verdict("darwin", "ssh host 'sed -i s/a/b/ f'"))
	assert.String(t, "", guard.Verdict("linux", "bash -c 'sed -i s/a/b/ f'"))
}
