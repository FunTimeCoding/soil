package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/service"
	"gorm.io/gorm"
)

type Daemon struct {
	database *gorm.DB
	service  *service.Service
}
