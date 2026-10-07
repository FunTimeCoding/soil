package mock_directory

import "github.com/funtimecoding/soil/pkg/directory/types/entry"

type Directory struct {
	entries   []*entry.Entry
	passwords map[string]string
	members   map[string]bool
	failure   error
}
