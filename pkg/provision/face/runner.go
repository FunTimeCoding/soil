package face

import "github.com/funtimecoding/soil/pkg/provision/types/trigger"

type Runner interface {
	Trigger(request trigger.Request) error
}
