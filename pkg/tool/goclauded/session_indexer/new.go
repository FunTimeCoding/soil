package session_indexer

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

func New(
	c *client.Client,
	collection string,
	sourceType string,
) *Indexer {
	return &Indexer{
		client:     c,
		collection: collection,
		sourceType: sourceType,
	}
}
