package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"time"
)

func (s *Service) SendKeys(
	identifier string,
	keys []string,
	interval time.Duration,
) error {
	if len(keys) == 0 {
		return validation.New("keys is required")
	}

	interval = max(interval, s.keyInterval)

	return s.attempt(
		func() error {
			for i, key := range keys {
				if i > 0 {
					time.Sleep(interval)
				}

				if e := s.client.SendKey(identifier, key); e != nil {
					return fmt.Errorf("step %d: %w", i+1, e)
				}
			}

			return nil
		},
	)
}
