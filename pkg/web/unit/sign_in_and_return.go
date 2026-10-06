package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func signInAndReturn(
	t *testing.T,
	c *client.Client,
	code string,
) *httptest.ResponseRecorder {
	t.Helper()
	in := httptest.NewRecorder()
	c.SignIn(in, httptest.NewRequest(http.MethodGet, constant.SignInPath, nil))
	l, e := url.Parse(in.Header().Get("Location"))
	assert.FatalOnError(t, e)
	back := httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/callback?state=%s&code=%s", l.Query().Get("state"), code),
		nil,
	)

	for _, cookie := range in.Result().Cookies() {
		back.AddCookie(cookie)
	}

	result := httptest.NewRecorder()
	c.Callback(result, back)

	return result
}
