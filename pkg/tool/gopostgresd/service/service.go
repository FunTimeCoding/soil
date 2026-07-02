package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/inventory"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Service struct {
	inventory *inventory.Inventory
	pools     map[string]*pgxpool.Pool
	sessions  sync.Map
	mutex     sync.Mutex
}
