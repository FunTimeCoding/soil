package service

import (
	"fmt"
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
)

func (s *Service) finish(
	k *watched_request.Key,
	subscriptions []subscription.Subscription,
	line string,
) []error {
	var result []error

	for _, v := range subscriptions {
		e := s.digests.Notify(
			v.Callsign,
			fmt.Sprintf("[%s] %s, stopped watching", v.Label(), line),
			goclaudedConstant.DeliveryQueued,
		)

		if refused := s.note(v.Callsign, e); refused != nil {
			result = append(result, refused)
		}
	}

	s.pending.Retain(k.String(), nil)

	if e := s.store.DeleteRequest(k); e != nil {
		result = append(result, e)
	}

	return result
}
