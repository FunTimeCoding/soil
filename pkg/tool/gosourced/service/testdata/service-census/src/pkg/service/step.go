package service

import (
	"example/pkg/model/receipt"
	"example/pkg/store"
)

func (s *Service) step(t *store.Store) error {
	t.SaveReceipt(&receipt.Receipt{})

	return nil
}
