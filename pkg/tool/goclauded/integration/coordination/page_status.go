package coordination

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"net/http"
	"testing"
)

func pageStatus(
	t *testing.T,
	s *base.Server,
	path string,
) int {
	t.Helper()
	response, e := http.Get(fmt.Sprintf("http://localhost:%d%s", s.Port, path))
	assert.FatalOnError(t, e)

	defer errors.PanicClose(response.Body)

	return response.StatusCode
}
