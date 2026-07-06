package service_client

import (
	"github.com/ctreminiom/go-atlassian/v2/jira/sm"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	user string,
	token string,
) *sm.Client {
	result, e := sm.New(nil, locator.New(host).String())
	errors.PanicOnError(e)
	result.Auth.SetBasicAuth(user, token)

	return result
}
