package jira

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/client"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/service_client"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	user string,
	token string,
	o ...Option,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(user, "user")
	errors.FatalOnEmpty(token, "token")
	result := &Client{
		context: context.Background(),
		client: client.New(
			jira.BasicAuthTransport{Username: user, Password: token},
			host,
		),
		basic:   basic.New(host, user, token),
		service: service_client.New(host, user, token),
		locator: locator.New(host).String(),
		user:    user,
	}

	for _, p := range o {
		p(result)
	}

	if len(result.closedStatus) == 0 {
		result.closedStatus = []string{constant.Closed}
	}

	// Verify token. Search V2 returns 200 even if token is invalid.
	_, e := result.User()
	errors.PanicOnError(e)

	return result
}
