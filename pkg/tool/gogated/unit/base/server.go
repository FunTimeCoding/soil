package base

import (
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gogated/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
	"gorm.io/gorm"
	"testing"
)

type Server struct {
	Reporter  *memory.Memory
	directory face.DirectoryClient
	litePath  string
	daemon    *daemon.Daemon
	database  *gorm.DB
	store     *store.Store
	service   *service.Service
	served    *model_context_server.Server
	t         *testing.T
}
