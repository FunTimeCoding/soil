package face

import "github.com/funtimecoding/soil/pkg/directory"

type Directory interface {
	Authenticate(account string, password string) (*directory.Entry, error)
	InGroup(account string) (bool, error)
}
