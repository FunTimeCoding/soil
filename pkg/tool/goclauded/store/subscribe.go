package store

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/subscription"
)

func (s *Store) Subscribe(
	name string,
	kinds []string,
) (*subscription.Subscription, error) {
	now := s.clock()
	var existing subscription.Subscription

	if e := s.database.Where(
		"name = ?",
		name,
	).First(&existing).Error; e == nil {
		existing.Kinds = join.Comma(kinds)
		existing.LastSeen = now

		if f := s.database.Save(&existing).Error; f != nil {
			return nil, f
		}

		return &existing, nil
	}

	result := subscription.New()
	result.Name = name
	result.Kinds = join.Comma(kinds)
	result.CreatedAt = now
	result.LastSeen = now

	if e := s.database.Create(result).Error; e != nil {
		return nil, e
	}

	return result, nil
}
