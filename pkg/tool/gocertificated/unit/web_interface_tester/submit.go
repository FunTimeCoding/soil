package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
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
		o.Server.Address(path),
		strings.NewReader(values.Encode()),
	)
	assert.FatalOnError(o.t, e)
	r.Header.Set(webConstant.ContentType, webConstant.FormEncoded)
	web.Bearer(r, generativeConstant.ModelContextTestToken)
	result, e := o.Server.Client().Do(r)
	assert.FatalOnError(o.t, e)

	return web.ReadString(result)
}
