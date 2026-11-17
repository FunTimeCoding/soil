package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
)

func (s *Server) Execute(statement string) {
	s.t.Helper()
	d := connection.New(logger.New(s.t.Context()), s.LitePath)
	defer errors.PanicClose(d)
	_, e := d.Exec(statement)
	assert.FatalOnError(s.t, e)
}
