package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/openbao/status"
)

type OpenbaoClient interface {
	Health() (*status.Status, error)
	List(path string) ([]string, error)
	Read(path string) (map[string]string, error)
	Put(
		path string,
		values map[string]string,
	) error
	Reachable(x context.Context) error
	Address() string
	Close()
}
