package memory_indexer

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/indexer"
)

type Indexer struct {
	*indexer.Indexer
	client *client.Client
}
