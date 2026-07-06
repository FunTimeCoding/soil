package github

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/google/go-github/v88/github"
	"golang.org/x/oauth2"
)

func New(token string) *Client {
	result := &Client{context: context.Background()}
	c, e := github.NewClient(
		github.WithHTTPClient(
			oauth2.NewClient(
				result.context,
				oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
			),
		),
	)
	errors.PanicOnError(e)
	result.client = c

	return result
}
