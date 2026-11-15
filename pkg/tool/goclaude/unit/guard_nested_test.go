package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"testing"
)

func TestSupplyChainInsideShellStringBlocked(t *testing.T) {
	assert.String(
		t,
		"npx is blocked (supply-chain guard) - it downloads and executes npm packages on demand",
		guard.BlockReason("darwin", "bash -c 'npx cowsay hi'"),
	)
	assert.String(
		t,
		"pip install is blocked (supply-chain guard) - no python dependencies may be installed on this system",
		guard.BlockReason("linux", `sh -c "pip install $package"`),
	)
	assert.String(
		t,
		"npx is blocked (supply-chain guard) - it downloads and executes npm packages on demand",
		guard.BlockReason("linux", "eval npx serve ."),
	)
}

func TestBsdFlagsInsideShellStringBlocked(t *testing.T) {
	assert.String(
		t,
		"sed on macOS is BSD sed and its flags (notably -i) differ from GNU sed - use gsed instead",
		guard.BlockReason("darwin", "zsh -c 'sed -i s/a/b/ file.txt'"),
	)
	assert.String(
		t,
		"xargs on macOS is BSD xargs and rejects the GNU flags -a and -d - redirect the file on stdin instead: xargs command < file",
		guard.BlockReason("darwin", "bash -lc 'xargs -a list.txt echo'"),
	)
}

func TestRemoteShellStringsAllowed(t *testing.T) {
	assert.String(
		t,
		"",
		guard.BlockReason("darwin", `kubectl exec pod -- sh -c "npx serve ."`),
	)
	assert.String(
		t,
		"",
		guard.BlockReason("darwin", "ssh host 'sed -i s/a/b/ f'"),
	)
	assert.String(
		t,
		"",
		guard.BlockReason("linux", "bash -c 'sed -i s/a/b/ f'"),
	)
}
