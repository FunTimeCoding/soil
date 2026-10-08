package conversation

import "github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/hit"

type Conversation struct {
	Session string
	Name    string
	Latest  string
	Count   int
	Hits    []*hit.Hit
}
