package delivery

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func fixedLines(groups map[string][]queue.Entry) []string {
	var roster []queue.Entry

	for _, kind := range []string{
		constant.QueueSessionAnnounce,
		constant.QueueSessionRelease,
		constant.QueueSessionComplete,
		constant.QueueSessionUpdate,
	} {
		roster = append(roster, groups[kind]...)
	}

	result := bodies(groups[constant.QueueReannounce], "%s")
	result = append(
		result,
		bodies(groups[constant.QueuePulse], constant.DeliveryPulse)...,
	)
	result = append(
		result,
		bodies(groups[constant.QueueTimeout], constant.DeliveryIdle)...,
	)
	result = append(
		result,
		section(
			constant.DeliverySessionActivity,
			bodies(roster, constant.DeliveryIndent),
		)...,
	)

	return append(
		result,
		section(
			constant.DeliveryNotifications,
			bodies(groups[constant.QueueNotification], constant.DeliveryIndent),
		)...,
	)
}
