package service

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"
	"time"
)

type Service struct {
	pool  *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
	clock func() time.Time
}
