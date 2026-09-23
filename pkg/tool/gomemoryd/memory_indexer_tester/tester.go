package memory_indexer_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/memory_indexer"
	"testing"
)

type Tester struct {
	*memory_indexer.Indexer
	t *testing.T
}
