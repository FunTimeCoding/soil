package store

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/subscription"
)

func (s *Store) SetSubscriptionPosition(
	name string,
	identifier uint,
) error {
	x := s.database.Model(subscription.Stub()).Where("name = ?", name).Updates(
		map[string]any{
			constant.EventIdentifierColumn: identifier,
			constant.LastSeenColumn:        s.clock(),
		},
	)

	if x.Error != nil {
		return x.Error
	}

	if x.RowsAffected == 0 {
		return not_found.New(constant.Subscription, name)
	}

	return nil
}
