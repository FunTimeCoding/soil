package mock_indexer

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/types/call"

type Indexer struct {
	Pushed    []*call.Push
	Deleted   []*call.Delete
	Documents map[string]string
}
