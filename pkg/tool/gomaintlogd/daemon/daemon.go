package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database  *gorm.DB
	service   *service.Service
	web       *web.Server
	closeOnce sync.Once
}
