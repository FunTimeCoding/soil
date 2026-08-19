package web_interface_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"net/url"
	"testing"
)

func (o *Tester) Submit(
	t *testing.T,
	path string,
	values url.Values,
) string {
	t.Helper()
	result, e := http.PostForm(
		fmt.Sprintf("http://127.0.0.1:%d%s", o.Server.Port(), path),
		values,
	)
	assert.FatalOnError(t, e)

	return web.ReadString(result)
}
