package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func sendForm(
	t *testing.T,
	s *model_context_server.Server,
	target string,
	values url.Values,
) (int, string) {
	t.Helper()
	var q *http.Request

	if values == nil {
		q = web.NewGet(s.Address(target))
	} else {
		q = web.NewPost(s.Address(target), values.Encode())
		q.Header.Set(constant.ContentType, constant.FormEncoded)
	}

	q.Header.Set(constant.ExtendedRequest, "true")
	r, e := s.Client().Do(q)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)

	return r.StatusCode, string(b)
}
