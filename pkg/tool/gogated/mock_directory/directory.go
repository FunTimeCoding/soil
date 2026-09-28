package mock_directory

import "github.com/funtimecoding/soil/pkg/directory"

type Directory struct {
	entries   []*directory.Entry
	passwords map[string]string
	members   map[string]bool
	failure   error
}
