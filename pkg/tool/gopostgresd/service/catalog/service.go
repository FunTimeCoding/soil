package catalog

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/peered"
)

type Service struct {
	pool *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
}
