package policy_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/allowlist"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/policy"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"path/filepath"
)

func (o *Tester) walk() *policy.Walk {
	o.t.Helper()
	rules, e := allowlist.Read(settingPath(o.Root))
	assert.FatalOnError(o.t, e)

	return policy.New(
		rules,
		[]string{
			filepath.Join(o.Root, constant.FixtureScratchDirectory),
			o.Scratchpad,
		},
		o.Root,
	)
}
