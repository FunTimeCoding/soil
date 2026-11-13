package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"net/url"
	"testing"
)

func statusCode(
	t *testing.T,
	s *model_context_server.Server,
	target string,
	form url.Values,
) int {
	t.Helper()
	c := s.Client()

	if form == nil {
		r, e := c.Get(s.Address(target))
		assert.FatalOnError(t, e)
		errors.PanicClose(r.Body)

		return r.StatusCode
	}

	r, e := c.PostForm(s.Address(target), form)
	assert.FatalOnError(t, e)
	errors.PanicClose(r.Body)

	return r.StatusCode
}
