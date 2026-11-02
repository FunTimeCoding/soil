package base

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"testing"
)

type Server struct {
	t        *testing.T
	database *sql.DB
	store    *store.Store
	daemon   *daemon.Daemon
	indexer  *mock_client.Client
	Reporter *memory.Memory
	*model_context_server.Server
}
