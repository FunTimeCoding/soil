package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"sync"
)

type Client struct {
	context      context.Context
	client       *gitlab.Client
	mutex        sync.Mutex
	user         *gitlab.User
	verbose      bool
	insecure     bool
	groups       []int64
	projects     []int64
	projectCache map[int64]*project.Project
	pool         *pool.Pool
}
