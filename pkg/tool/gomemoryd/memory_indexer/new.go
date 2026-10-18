package memory_indexer

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/indexer"
)

func New(c *client.ClientWithResponses) *Indexer {
	return &Indexer{
		Indexer: indexer.New(c, constant.MemorySourceType),
		client:  c,
	}
}
