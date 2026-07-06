package event

import (
	"github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"k8s.io/api/events/v1"
)

func New(
	v *v1.Event,
	cluster string,
) *Event {
	return &Event{
		MonitorIdentifier: constant.GoKevt.StringIdentifier(
			v.Name,
		),
		Cluster:       cluster,
		Namespace:     v.Namespace,
		RegardingKind: v.Regarding.Kind,
		Name:          v.Name,
		Type:          v.Type,
		Reason:        v.Reason,
		Create:        &v.CreationTimestamp.Time,
		Note:          v.Note,
		Raw:           v,
	}
}
