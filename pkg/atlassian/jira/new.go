package jira

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/client"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/service_client"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
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
	p := stall_client.New().Build()
	result := &Client{
		context: context.Background(),
		client: client.New(
			jira.BasicAuthTransport{Username: user, Password: token},
			host,
			p.Client().Transport,
		),
		basic:   basic.New(locator.New(host), user, token),
		service: service_client.New(host, user, token, p.Client()),
		pool:    p,
		locator: locator.New(host).String(),
		user:    user,
	}

	for _, p := range o {
		p(result)
	}

	if len(result.closedStatus) == 0 {
		result.closedStatus = []string{constant.JiraClosed}
	}

	_, e := result.User()
	errors.PanicOnError(e)

	return result
}
