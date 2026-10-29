package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
	"gorm.io/gorm"
)

type Server struct {
	Database *gorm.DB
	Store    *store.Store
	Service  *service.Service
	Reporter *memory.Memory
	Web      *model_context_server.Server
}
