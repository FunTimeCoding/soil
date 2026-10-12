package face

import "github.com/funtimecoding/soil/pkg/openbao/status"

type SecretSource interface {
	Health() (*status.Status, error)
	List(path string) ([]string, error)
	Read(path string) (map[string]string, error)
	Put(
		path string,
		values map[string]string,
	) error
}
