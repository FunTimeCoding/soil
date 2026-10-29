package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage"
)

type UsageClient interface {
	Fetch() (*usage.Usage, error)
	Reachable(x context.Context) error
	Close()
}
