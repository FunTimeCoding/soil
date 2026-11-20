package lint_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
)

func (o *Tester) Report(
	path string,
	concerns bool,
	c []*concern.Concern,
	fixed string,
	l *file_report.Report,
) {
	o.t.Helper()
	assert.String(o.t, path, l.FilePath)
	assert.Boolean(o.t, concerns, l.HasConcerns())
	assert.Any(o.t, c, l.Concerns)
	assert.String(o.t, fixed, l.Fixed)
}
