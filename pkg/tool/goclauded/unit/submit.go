package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func submit(
	t *testing.T,
	s *base.Server,
	path string,
	values url.Values,
	extended bool,
) string {
	t.Helper()
	q, e := http.NewRequest(
		http.MethodPost,
		s.Address(path),
		strings.NewReader(values.Encode()),
	)
	assert.FatalOnError(t, e)
	q.Header.Set(webConstant.ContentType, webConstant.FormEncoded)
	web.Bearer(q, generativeConstant.ModelContextTestToken)

	if extended {
		q.Header.Set(webConstant.ExtendedRequest, "true")
	}

	r, f := s.Server.Client().Do(q)
	assert.FatalOnError(t, f)

	return web.ReadString(r)
}
