package client

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	t jira.BasicAuthTransport,
	host string,
) *jira.Client {
	result, e := jira.NewClient(t.Client(), locator.New(host).String())
	errors.PanicOnError(e)

	return result
}
