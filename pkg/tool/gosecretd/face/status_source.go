package face

import "github.com/funtimecoding/soil/pkg/openbao/status"

type StatusSource interface {
	Status() (*status.Status, error)
}
