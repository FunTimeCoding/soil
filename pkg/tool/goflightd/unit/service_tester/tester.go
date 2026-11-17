package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/service"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/mock_system"
	"gorm.io/gorm"
	"testing"
)

type Tester struct {
	Store    *store.Store
	System   *mock_system.System
	Reporter *memory.Memory
	mapper   *gorm.DB
	service  *service.Service
	t        *testing.T
}
