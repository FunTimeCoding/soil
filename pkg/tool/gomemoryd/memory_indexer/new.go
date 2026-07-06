package memory_indexer

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

func New(c *client.Client) *Indexer {
	return &Indexer{client: c}
}
