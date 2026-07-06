package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type Client struct {
	context      context.Context
	client       *gitlab.Client
	user         *gitlab.User
	verbose      bool
	groups       []int64
	projects     []int64
	projectCache map[int64]*project.Project
}
