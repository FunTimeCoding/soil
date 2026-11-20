package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"go/format"
)

func (o *Tester) Formatted(content string) {
	o.t.Helper()
	formatted, e := format.Source([]byte(content))
	assert.FatalOnError(o.t, e)
	assert.String(o.t, string(formatted), content)
}
