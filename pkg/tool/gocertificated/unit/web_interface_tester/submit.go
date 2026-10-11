package web_interface_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/url"
	"strings"
)

func (o *Tester) Submit(
	path string,
	values url.Values,
) string {
	o.t.Helper()
	r, e := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://127.0.0.1:%d%s", o.Server.Port, path),
		strings.NewReader(values.Encode()),
	)
	assert.FatalOnError(o.t, e)
	r.Header.Set(constant.ContentType, constant.FormEncoded)
	r.AddCookie(o.Server.Authorization.SubjectCookie("tester"))
	result, e := http.DefaultClient.Do(r)
	assert.FatalOnError(o.t, e)

	return web.ReadString(result)
}
