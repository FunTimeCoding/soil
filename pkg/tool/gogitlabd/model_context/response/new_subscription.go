package response

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"

func NewSubscription(v subscription.Subscription) *Subscription {
	return &Subscription{
		Reference:    v.Reference,
		Alias:        v.Alias,
		Instance:     v.Instance,
		Project:      v.Project,
		MergeRequest: v.Request,
		Delivery:     string(v.DeliveryMode()),
	}
}
