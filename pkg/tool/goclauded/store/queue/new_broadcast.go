package queue

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func NewBroadcast(
	kind string,
	body string,
) *Entry {
	return &Entry{Kind: kind, Body: body, Delivery: constant.DeliveryQueued}
}
