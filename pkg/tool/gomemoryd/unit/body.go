package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"testing"
)

func body(
	t *testing.T,
	s *model_context_server.Server,
	target string,
	extended bool,
) string {
	t.Helper()
	q := web.NewGet(s.Address(target))

	if extended {
		q.Header.Set(constant.ExtendedRequest, "true")
	}

	r, e := s.Client().Do(q)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)

	return string(b)
}
