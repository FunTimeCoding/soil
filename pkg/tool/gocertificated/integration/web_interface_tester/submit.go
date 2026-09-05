package web_interface_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func (o *Tester) Submit(
	t *testing.T,
	path string,
	values url.Values,
) string {
	t.Helper()
	r, e := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://127.0.0.1:%d%s", o.Server.Port(), path),
		strings.NewReader(values.Encode()),
	)
	assert.FatalOnError(t, e)
	r.Header.Set(constant.ContentType, constant.FormEncoded)
	r.AddCookie(o.Server.Authorization.SubjectCookie("tester"))
	result, e := http.DefaultClient.Do(r)
	assert.FatalOnError(t, e)

	return web.ReadString(result)
}
