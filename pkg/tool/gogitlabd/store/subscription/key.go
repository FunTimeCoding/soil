package subscription

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"

func (s Subscription) Key() *watched_request.Key {
	return watched_request.New(s.Instance, s.Project, s.Request)
}
