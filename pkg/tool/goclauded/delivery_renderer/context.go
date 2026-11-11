package delivery_renderer

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (r *Renderer) Context(entries []queue.Entry) string {
	groups := group(entries)
	candidates := groups[constant.QueueMessage]
	lines := fixedLines(groups)
	lines = append(
		lines,
		memoryLines(
			groups,
			constant.DeliveryBudget-size(lines)-r.messageMinimum(candidates),
		)...,
	)
	lines = append(
		lines,
		r.messageLines(candidates, constant.DeliveryBudget-size(lines))...,
	)

	return join.NewLine(lines)
}
