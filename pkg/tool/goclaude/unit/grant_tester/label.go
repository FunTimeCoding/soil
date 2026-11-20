package grant_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
)

func (r *Tester) Label(repository string) {
	r.t.Helper()
	assert.FatalOnError(
		r.t,
		r.Mock.AddLabel(r.Session,constant.GitWriteLabel, repository, ""),
	)
}
