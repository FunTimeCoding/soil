package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/vulnerability"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web"
	"github.com/funtimecoding/soil/pkg/vulnerability/face"
	"github.com/funtimecoding/soil/pkg/worker"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database         *gorm.DB
	nationalDatabase face.NationalDatabase
	nationalGuard    *peer.Guard
	securityTracker  face.SecurityTracker
	trackerGuard     *peer.Guard
	extendedTracker  face.SecurityTracker
	extendedGuard    *peer.Guard
	vulnerability    *vulnerability.Service
	debian           *debian.Service
	flag             *flag.Service
	system           *system.Service
	web              *web.Server
	clock            *worker.Worker
	closeOnce        sync.Once
}
