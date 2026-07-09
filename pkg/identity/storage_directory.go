package identity

import "github.com/funtimecoding/soil/pkg/system"

func (t *Tool) StorageDirectory(create bool) string {
	return system.StorageDirectory(t.name, create)
}
