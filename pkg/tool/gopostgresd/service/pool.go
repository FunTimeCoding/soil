package service

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (s *Service) Pool(instance string) (*pgxpool.Pool, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if p, okay := s.pools[instance]; okay {
		return p, nil
	}

	i, okay := s.Instance(instance)

	if !okay {
		return nil, fmt.Errorf("unknown instance: %s", instance)
	}

	p, e := pgxpool.New(context.Background(), i.Locator())

	if e != nil {
		return nil, fmt.Errorf("connect to %s: %w", instance, e)
	}

	s.pools[instance] = p

	return p, nil
}
