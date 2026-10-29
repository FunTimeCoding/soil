package face

import "github.com/funtimecoding/soil/pkg/provision/types/update"

type Provision interface {
	Runner
	Runs
	Sync() (*update.Result, error)
}
