package base

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"testing"
)

type Server struct {
	t        *testing.T
	database *sql.DB
	store    *store.Store
	service  *service.Service
	indexer  *mock_client.Client
	*model_context_server.Server
}
