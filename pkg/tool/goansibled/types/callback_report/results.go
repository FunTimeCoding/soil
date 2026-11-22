package callback_report

import (
	"github.com/funtimecoding/soil/pkg/provision/types/host_result"
	"slices"
	"strings"
)

func (r *Report) Results() []*host_result.Result {
	var result []*host_result.Result

	for _, p := range r.Plays {
		for _, t := range p.Tasks {
			for host, o := range t.Hosts {
				result = append(result, hostResult(host, o))
			}
		}
	}

	slices.SortFunc(
		result,
		func(a, b *host_result.Result) int {
			return strings.Compare(a.Host, b.Host)
		},
	)

	return result
}
