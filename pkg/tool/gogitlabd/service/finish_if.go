package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/change"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
)

func (s *Service) finishIf(
	k *watched_request.Key,
	subscriptions []subscription.Subscription,
	v *snapshot.Snapshot,
	actor string,
) []error {
	if !v.Finished() {
		return nil
	}

	return s.finish(k, subscriptions, change.StateLine(v.State, actor))
}
