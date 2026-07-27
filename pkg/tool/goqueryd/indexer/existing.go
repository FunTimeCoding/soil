package indexer

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/connect"

func (i *Indexer) Existing(collection string) map[string]string {
	return connect.Existing(i.client, collection)
}
