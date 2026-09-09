package face

import "github.com/funtimecoding/soil/pkg/nextcloud/usage"

type UsageSource interface {
	Fetch() (*usage.Usage, error)
}
