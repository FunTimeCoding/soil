package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Service struct {
	inventory *inventory.Inventory
	pools     map[string]*pgxpool.Pool
	sessions  sync.Map
	mutex     sync.Mutex
}
