package mock_client

import "github.com/funtimecoding/soil/pkg/directory/types/entry"

type Client struct {
	Failure   error
	base      string
	records   map[string]map[string][]string
	order     []string
	entries   []*entry.Entry
	passwords map[string]string
	members   map[string]bool
}
