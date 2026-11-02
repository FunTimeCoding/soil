package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_entry"
	"slices"
)

func (s *Service) Requests() []*request_entry.Entry {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	var result []*request_entry.Entry

	for _, i := range s.inventory.Instances {
		result = append(result, s.requests[i.Name]...)
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
