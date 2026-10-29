package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/web"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func postGenerate(
	t *testing.T,
	v *model_context_server.Server,
	names ...string,
) (int, string) {
	t.Helper()
	q, e := http.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		v.Address("/generate"),
		strings.NewReader(url.Values{"fileNames": names}.Encode()),
	)
	assert.FatalOnError(t, e)
	q.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	web.Bearer(q, constant.ModelContextTestToken)
	r, f := v.Halting().Do(q)
	assert.FatalOnError(t, f)
	defer errors.LogClose(r.Body)
	b, g := io.ReadAll(r.Body)
	assert.FatalOnError(t, g)

	return r.StatusCode, string(b)
}
