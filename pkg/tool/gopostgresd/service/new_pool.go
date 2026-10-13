package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/jackc/pgx/v5/pgxpool"
)

func newPool(i *instance.Instance) (*pgxpool.Pool, error) {
	result, e := pgxpool.New(context.Background(), i.Locator())

	if e != nil {
		return nil, fmt.Errorf("connect to %s: %w", i.Name, e)
	}

	return result, nil
}
