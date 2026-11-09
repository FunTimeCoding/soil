package backup

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/store/backup"
	"time"
)

type Service struct {
	inventory *instance.Inventory
	pool      *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
	store     *backup.Store
	dump      face.Dump
	clock     func() time.Time
}
