package store

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"

func (s *Store) Abandon(present []string) ([]subscription.Subscription, error) {
	query := s.database

	if len(present) > 0 {
		query = query.Where("callsign NOT IN ?", present)
	}

	var result []subscription.Subscription

	if e := query.Find(&result).Error; e != nil {
		return nil, e
	}

	for _, v := range result {
		if e := s.database.Delete(
			subscription.Stub(),
			v.Identifier,
		).Error; e != nil {
			return nil, e
		}
	}

	return result, nil
}
