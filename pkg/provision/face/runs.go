package face

import "github.com/funtimecoding/soil/pkg/provision/model/run"

type Runs interface {
	Find(identifier uint) (*run.Run, error)
	Recent(limit int) ([]run.Run, error)
}
