package daemon

import (
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/web"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/worker"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database  *gorm.DB
	client    face.NetBoxClient
	guard     *peer.Guard
	service   *service.Service
	web       *web.Server
	clock     *worker.Worker
	closeOnce sync.Once
}
