package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/mock_system"
	"gorm.io/gorm"
	"testing"
	"time"
)

type Server struct {
	System   *mock_system.System
	Reporter *memory.Memory
	clock    func() time.Time
	litePath string
	daemon   *daemon.Daemon
	mapper   *gorm.DB
	served   *model_context_server.Server
	t        *testing.T
}
