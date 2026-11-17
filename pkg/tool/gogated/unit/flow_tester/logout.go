package flow_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"net/http"
)

func (o *Tester) Logout(authenticationCookie *http.Cookie) {
	o.t.Helper()
	request, e := http.NewRequest(
		http.MethodPost,
		join.Empty(o.server.BaseLocator(), constant.LogoutPath),
		nil,
	)
	assert.FatalOnError(o.t, e)

	if authenticationCookie != nil {
		request.AddCookie(authenticationCookie)
	}

	response, e := o.client.Do(request)
	assert.FatalOnError(o.t, e)
	errors.PanicClose(response.Body)
	assert.Integer(o.t, http.StatusOK, response.StatusCode)
}
