package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"io"
	"testing"
)

func body(
	t *testing.T,
	s *model_context_server.Server,
	target string,
) string {
	t.Helper()
	r, e := s.Client().Get(s.Address(target))
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)

	return string(b)
}
