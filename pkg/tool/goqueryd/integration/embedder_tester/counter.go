package embedder_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/embed_counter"
)

func (o *Tester) Counter() *embed_counter.Counter {
	o.t.Helper()
	result, e := embed_counter.NewEnvironment()
	assert.FatalOnError(o.t, e)
	o.t.Cleanup(func() { errors.PanicClose(result) })

	return result
}
