package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func New(
	host string,
	token string,
	o ...Option,
) *Client {
	result := &Client{
		context:      context.Background(),
		projectCache: make(map[int64]*project.Project),
	}

	for _, p := range o {
		p(result)
	}

	var p []gitlab.ClientOptionFunc

	if host != "" {
		p = append(
			p,
			gitlab.WithBaseURL(
				locator.New(host).Base(constant.Base).String(),
			),
		)
	}

	client, e := gitlab.NewClient(token, p...)
	errors.PanicOnError(e)
	result.client = client
	result.user = result.MustCurrentUser()

	return result
}
