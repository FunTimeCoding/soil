package worker

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_entry"
	"slices"
)

func (w *Worker) Requests() []*request_entry.Entry {
	w.mutex.RLock()
	defer w.mutex.RUnlock()
	var result []*request_entry.Entry

	for _, i := range w.service.Inventory().Instances {
		result = append(result, w.requests[i.Name]...)
	}

	slices.SortStableFunc(
		result,
		func(
			a *request_entry.Entry,
			b *request_entry.Entry,
		) int {
			if a.Request.Create == nil || b.Request.Create == nil {
				return 0
			}

			return b.Request.Create.Compare(*a.Request.Create)
		},
	)

	if len(result) > constant.RequestLimit {
		result = result[:constant.RequestLimit]
	}

	return result
}
