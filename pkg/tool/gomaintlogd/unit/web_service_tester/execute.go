package web_service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
)

func (o *Tester) Execute(statement string) {
	o.t.Helper()
	d := connection.New(logger.New(o.t.Context()), o.server.LitePath)
	defer errors.PanicClose(d)
	_, e := d.Exec(statement)
	assert.FatalOnError(o.t, e)
}
