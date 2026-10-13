package service

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (s *Service) Pool() *pool.Pool[instance.Instance, *instance.Instance, *pgxpool.Pool] {
	return s.pool
}
