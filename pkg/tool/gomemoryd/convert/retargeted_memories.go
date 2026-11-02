package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/retarget_report"
)

func RetargetedMemories(
	reports []*retarget_report.Report,
) []server.RetargetedMemory {
	result := []server.RetargetedMemory{}

	for _, r := range reports {
		changes := []server.RetargetChange{}

		for _, c := range r.Changes {
			changes = append(
				changes,
				server.RetargetChange{
					Line:   c.Line,
					Before: c.Before,
					After:  c.After,
					Reason: c.Reason,
				},
			)
		}

		result = append(
			result,
			server.RetargetedMemory{
				Identifier: r.Identifier,
				Name:       r.Name,
				Changes:    changes,
			},
		)
	}

	return result
}
