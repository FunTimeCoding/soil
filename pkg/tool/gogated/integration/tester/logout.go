package tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"net/http"
	"testing"
)

func (o *Tester) Logout(
	t *testing.T,
	authenticationCookie *http.Cookie,
) {
	t.Helper()
	request, e := http.NewRequest(
		http.MethodPost,
		join.Empty(o.server.BaseLocator(), constant.LogoutPath),
		nil,
	)
	assert.FatalOnError(t, e)

	if authenticationCookie != nil {
		request.AddCookie(authenticationCookie)
	}

	response, e := o.client.Do(request)
	assert.FatalOnError(t, e)
	errors.PanicClose(response.Body)
	assert.Integer(t, http.StatusOK, response.StatusCode)
}
