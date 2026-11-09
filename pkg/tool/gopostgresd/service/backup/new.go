package backup

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/store/backup"
	"time"
)

func New(
	i *instance.Inventory,
	p *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered],
	s *backup.Store,
	d face.Dump,
	clock func() time.Time,
) *Service {
	return &Service{inventory: i, pool: p, store: s, dump: d, clock: clock}
}
