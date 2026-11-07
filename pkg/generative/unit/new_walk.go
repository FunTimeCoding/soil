package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/allowlist"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/policy"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"path/filepath"
	"testing"
)

func newWalk(
	t *testing.T,
	root string,
	scratchpad string,
) *policy.Walk {
	t.Helper()
	rules, e := allowlist.Read(settingPath(root))
	assert.FatalOnError(t, e)

	return policy.New(
		rules,
		[]string{
			filepath.Join(root, constant.FixtureScratchDirectory),
			scratchpad,
		},
		root,
	)
}
