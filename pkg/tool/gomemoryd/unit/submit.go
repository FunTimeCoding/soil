package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"io"
	"net/url"
	"testing"
)

func submit(
	t *testing.T,
	s *model_context_server.Server,
	target string,
	form url.Values,
) string {
	t.Helper()
	r, e := s.Client().PostForm(s.Address(target), form)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)

	return string(b)
}
