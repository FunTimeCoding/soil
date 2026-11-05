package github

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/google/go-github/v93/github"
	"golang.org/x/oauth2"
)

func build(
	token string,
	options ...github.ClientOptionsFunc,
) *Client {
	result := &Client{
		context: context.Background(),
		pool:    stall_client.New().Build(),
	}
	c, e := github.NewClient(
		append(
			[]github.ClientOptionsFunc{
				github.WithHTTPClient(
					oauth2.NewClient(
						context.WithValue(
							result.context,
							oauth2.HTTPClient,
							result.pool.Client(),
						),
						oauth2.StaticTokenSource(
							&oauth2.Token{AccessToken: token},
						),
					),
				),
			},
			options...,
		)...,
	)
	errors.PanicOnError(e)
	result.client = c

	return result
}
