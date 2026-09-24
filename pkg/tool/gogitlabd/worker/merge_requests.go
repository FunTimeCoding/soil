package worker

import (
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"slices"
)

func MergeRequests(
	assigned []*merge_request.Request,
	reviewing []*merge_request.Request,
) []*merge_request.Request {
	seen := make(map[requestKey]bool)
	result := make([]*merge_request.Request, 0, len(assigned)+len(reviewing))

	for _, r := range slices.Concat(assigned, reviewing) {
		k := requestKey{r.Project, r.Identifier}

		if seen[k] {
			continue
		}

		seen[k] = true
		result = append(result, r)
	}

	slices.SortFunc(
		result,
		func(
			a *merge_request.Request,
			b *merge_request.Request,
		) int {
			if a.Create == nil || b.Create == nil {
				return 0
			}

			return b.Create.Compare(*a.Create)
		},
	)

	if len(result) > constant.RequestLimit {
		result = result[:constant.RequestLimit]
	}

	return result
}
