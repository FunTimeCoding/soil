package delivery

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (d *Delivery) Context(entries []queue.Entry) string {
	groups := group(entries)
	candidates := groups[constant.QueueMessage]
	lines := fixedLines(groups)
	lines = append(
		lines,
		memoryLines(
			groups,
			constant.DeliveryBudget-size(lines)-d.messageMinimum(candidates),
		)...,
	)
	lines = append(
		lines,
		d.messageLines(candidates, constant.DeliveryBudget-size(lines))...,
	)

	return join.NewLine(lines)
}
