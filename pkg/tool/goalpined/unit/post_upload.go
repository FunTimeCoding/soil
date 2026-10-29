package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"testing"
)

func postUpload(
	t *testing.T,
	v *model_context_server.Server,
	path string,
	body []byte,
) int {
	t.Helper()
	q, e := http.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		v.Address(webConstant.RootPath),
		bytes.NewReader(body),
	)
	assert.FatalOnError(t, e)
	q.URL.Opaque = path
	q.Header.Set(webConstant.ContentType, webConstant.Octets)
	web.Bearer(q, constant.ModelContextTestToken)
	r, f := v.Halting().Do(q)
	assert.FatalOnError(t, f)
	defer errors.LogClose(r.Body)

	return r.StatusCode
}
