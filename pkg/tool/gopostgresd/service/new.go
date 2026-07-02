package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/inventory"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(i *inventory.Inventory) *Service {
	return &Service{
		inventory: i,
		pools:     make(map[string]*pgxpool.Pool),
	}
}
