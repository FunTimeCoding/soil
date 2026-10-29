package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/lease"
	atlas "github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	database      *gorm.DB
	store         *store.Store
	authorization *client.Client
	netbox        netbox.Client
	netboxGuard   *peer.Guard
	peers         []face.Closer
	guards        []*peer.Guard
	collectors    []atlas.Collector
	lease         *lease.Collector
	retention     time.Duration
}
