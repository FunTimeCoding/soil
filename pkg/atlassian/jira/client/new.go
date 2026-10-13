package client

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func New(
	t jira.BasicAuthTransport,
	host string,
	r http.RoundTripper,
) *jira.Client {
	t.Transport = r
	result, e := jira.NewClient(t.Client(), locator.New(host).String())
	errors.PanicOnError(e)

	return result
}
