package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
)

func (s *Store) Find(
	callsign string,
	k *watched_request.Key,
) (*subscription.Subscription, bool, error) {
	all, e := s.ByRequest(k)

	if e != nil {
		return nil, false, e
	}

	for _, v := range all {
		if v.Callsign == callsign {
			return &v, true, nil
		}
	}

	return nil, false, nil
}
