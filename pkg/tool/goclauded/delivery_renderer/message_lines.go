package delivery_renderer

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (r *Renderer) messageLines(
	candidates []queue.Entry,
	room int,
) []string {
	whole := section(
		constant.DeliveryMessages,
		bodies(candidates, constant.DeliveryIndent),
	)

	if size(whole) <= room {
		return whole
	}

	result := []string{
		constant.DeliveryMessages,
		fmt.Sprintf(constant.DeliveryIndent, constant.DeliveryExplanation),
	}
	remaining := room - size(result)

	for i, e := range candidates {
		line := r.place(e, remaining-r.reserve(candidates[i+1:]))
		result = append(result, line)
		remaining -= lineSize(line)
	}

	return result
}
