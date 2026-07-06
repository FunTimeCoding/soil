package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_indexer"
	"testing"
)

type Server struct {
	t       *testing.T
	store   *store.Store
	indexer *mock_indexer.Indexer
	server  *model_context_server.Server
}
