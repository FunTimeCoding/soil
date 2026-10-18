package indexer

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

type Indexer struct {
	client     *client.ClientWithResponses
	sourceType string
}
