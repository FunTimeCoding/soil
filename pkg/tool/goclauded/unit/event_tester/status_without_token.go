package event_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
)

func (o *Tester) StatusWithoutToken(s *base.Server) int {
	o.t.Helper()
	r, e := s.Server.Client().Get(
		s.Address("/event-stream?subscriber=consumer"),
	)
	assert.FatalOnError(o.t, e)
	defer errors.PanicClose(r.Body)

	return r.StatusCode
}
