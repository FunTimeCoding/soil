package service_client

import (
	"github.com/ctreminiom/go-atlassian/v2/jira/sm"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func New(
	host string,
	user string,
	token string,
	h *http.Client,
) *sm.Client {
	result, e := sm.New(h, locator.New(host).String())
	errors.PanicOnError(e)
	result.Auth.SetBasicAuth(user, token)

	return result
}
