package fixture

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func StreamStatusWithoutToken(
	t *testing.T,
	s *base.Server,
) int {
	t.Helper()
	r, e := s.Server.Client().Get(
		s.Address("/event-stream?subscriber=consumer"),
	)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)

	return r.StatusCode
}
