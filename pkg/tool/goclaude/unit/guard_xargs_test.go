package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"testing"
)

func TestVerdictBlocksLocalXargsUnsupportedFlag(t *testing.T) {
	assert.String(
		t,
		"xargs on macOS is BSD xargs and rejects the GNU flags -a and -d - redirect the file on stdin instead: xargs command < file",
		guard.Verdict("darwin", "xargs -a list.txt git checkout --"),
	)
	assert.StringContains(
		t,
		"redirect the file on stdin",
		guard.Verdict("darwin", "xargs --arg-file=list.txt rm"),
	)
	assert.StringContains(
		t,
		"redirect the file on stdin",
		guard.Verdict("darwin", "xargs -d, echo"),
	)
	assert.StringContains(
		t,
		"redirect the file on stdin",
		guard.Verdict("darwin", "xargs --delimiter=, echo"),
	)
	assert.StringContains(
		t,
		"redirect the file on stdin",
		guard.Verdict("darwin", "/usr/bin/xargs -a list.txt echo"),
	)
	assert.StringContains(
		t,
		"redirect the file on stdin",
		guard.Verdict("darwin", "cat other && xargs -a list.txt echo"),
	)
}

func TestVerdictAllowsRemoteXargs(t *testing.T) {
	assert.String(
		t,
		"",
		guard.Verdict("darwin", `ssh root@10.0.0.2 "xargs -a list.txt rm"`),
	)
	assert.String(
		t,
		"",
		guard.Verdict(
			"darwin",
			`podman run --rm alpine sh -c "xargs -a list.txt echo"`,
		),
	)
	assert.String(
		t,
		"",
		guard.Verdict("darwin", `kubectl exec pod -- sh -c "xargs -d, echo"`),
	)
}

func TestVerdictAllowsSupportedXargsFlags(t *testing.T) {
	assert.String(
		t,
		"",
		guard.Verdict("darwin", "xargs git checkout -- < list.txt"),
	)
	assert.String(
		t,
		"",
		guard.Verdict("darwin", "find . -print0 | xargs -0 rm"),
	)
	assert.String(
		t,
		"",
		guard.Verdict("darwin", "grep -l x * | xargs -I{} echo {}"),
	)
	assert.String(
		t,
		"",
		guard.Verdict("darwin", "cat list | xargs -n1 -P4 echo"),
	)
	assert.String(t, "", guard.Verdict("darwin", "cat list | xargs -r echo"))
	assert.String(t, "", guard.Verdict("darwin", "cat list | xargs -L1 echo"))
}

func TestVerdictAllowsXargsOnLinux(t *testing.T) {
	assert.String(t, "", guard.Verdict("linux", "xargs -a list.txt echo"))
}

func TestVerdictXargsParseFallback(t *testing.T) {
	assert.StringContains(
		t,
		"redirect the file on stdin",
		guard.Verdict("darwin", "xargs -a 'list.txt echo"),
	)
}
