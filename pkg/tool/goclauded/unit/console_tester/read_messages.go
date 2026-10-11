package console_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude"
)

func (o *Tester) ReadMessages(identifiers ...int) string {
	o.t.Helper()
	result, e := goclaude.ReadMessages(o.client, identifiers)
	assert.FatalOnError(o.t, e)

	return result
}
