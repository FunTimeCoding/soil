package report

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
	"slices"
	"strconv"
)

func receiverDestinations(pairs []*entry.Entry) map[string][]string {
	result := map[string][]string{}

	for _, p := range pairs {
		if p.From.Receiver == "" || p.To.Receiver == "" {
			continue
		}

		from := join.Empty(
			strconv.Quote(p.From.Path),
			constant.ReferenceSeparator,
			p.From.Receiver,
		)
		to := join.Empty(
			strconv.Quote(p.To.Path),
			constant.ReferenceSeparator,
			p.To.Receiver,
		)

		if !slices.Contains(result[from], to) {
			result[from] = append(result[from], to)
		}
	}

	return result
}
