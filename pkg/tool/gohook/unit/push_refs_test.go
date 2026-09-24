package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git/changed"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"testing"
)

func TestPushRefsUpdate(t *testing.T) {
	r := changed.NewPushRefs(
		"",
		"refs/heads/main abc123 refs/heads/main def456\n",
	)
	assert.False(t, r.None)
	assert.False(t, r.All)
	assert.String(t, "def456..abc123", r.String())
}

func TestPushRefsDeletionOnly(t *testing.T) {
	r := changed.NewPushRefs(
		"",
		join.Space("(delete)", constant.ZeroHash, "refs/heads/old", "abc123\n"),
	)
	assert.True(t, r.None)
	assert.String(t, "nothing pushed", r.String())
	assert.Count(t, 0, r.Files(""))
}

func TestPushRefsNewBranchFallsBackToUpstream(t *testing.T) {
	r := changed.NewPushRefs(
		t.TempDir(),
		join.Space(
			"refs/heads/feature",
			"abc123",
			"refs/heads/feature",
			constant.ZeroHash,
		),
	)
	assert.False(t, r.None)
	assert.True(t, r.All)
}

func TestPushRefsSkipsDeletionBeforeUpdate(t *testing.T) {
	r := changed.NewPushRefs(
		"",
		join.Empty(
			join.Space(
				"(delete)",
				constant.ZeroHash,
				"refs/heads/old",
				"abc123\n",
			),
			"refs/heads/main 111 refs/heads/main 222\n",
		),
	)
	assert.String(t, "222..111", r.String())
}

func TestPushRefsEmptyInputUsesUpstream(t *testing.T) {
	assert.True(t, changed.NewPushRefs(t.TempDir(), "").All)
}
