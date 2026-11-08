package store

import (
	"example/pkg/model/order"
	"example/pkg/model/receipt"
	"example/pkg/store/result"
	"gorm.io/gorm"
)

type Store struct {
	mapper *gorm.DB
}

func (s *Store) Order(name string) *order.Order {
	var result order.Order
	s.mapper.Where("name = ?", name).First(&result)

	return &result
}

func (s *Store) Tally() *result.Tally {
	var t result.Tally
	s.mapper.Model(&order.Order{}).Scan(&t)

	return &t
}

func (s *Store) SaveReceipt(r *receipt.Receipt) {
	s.mapper.Save(r)
}

func (s *Store) Keep(r *receipt.Receipt) {
	s.SaveReceipt(r)
}

func (s *Store) Atomically(step func(*Store) error) error {
	return step(s)
}
