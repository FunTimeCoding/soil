package memory_indexer

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/connect"

func (i *Indexer) Existing() map[string]string {
	return connect.Existing(i.client, "memories")
}
