package face

import "github.com/funtimecoding/soil/pkg/directory/types/entry"

type Directory interface {
	Authenticate(account string, password string) (*entry.Entry, error)
	InGroup(account string) (bool, error)
}
