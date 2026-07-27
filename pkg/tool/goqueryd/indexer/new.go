package indexer

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

func New(c *client.Client, sourceType string) *Indexer {
	return &Indexer{client: c, sourceType: sourceType}
}
