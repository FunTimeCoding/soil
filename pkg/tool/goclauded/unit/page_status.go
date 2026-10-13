package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func pageStatus(
	t *testing.T,
	s *base.Server,
	path string,
) int {
	t.Helper()
	response, e := s.Server.Client().Get(s.Address(path))
	assert.FatalOnError(t, e)

	defer errors.PanicClose(response.Body)

	return response.StatusCode
}
