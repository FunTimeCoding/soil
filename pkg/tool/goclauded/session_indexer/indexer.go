package session_indexer

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

type Indexer struct {
	client     *client.Client
	collection string
	sourceType string
}
