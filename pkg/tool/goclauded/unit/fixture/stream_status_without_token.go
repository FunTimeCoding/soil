package fixture

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"net/http"
	"testing"
)

func StreamStatusWithoutToken(
	t *testing.T,
	s *base.Server,
) int {
	t.Helper()
	r, e := http.Get(
		fmt.Sprintf(
			"http://localhost:%d/event-stream?subscriber=consumer",
			s.Port,
		),
	)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)

	return r.StatusCode
}
