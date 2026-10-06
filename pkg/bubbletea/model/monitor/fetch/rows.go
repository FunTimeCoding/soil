package fetch

import (
	"charm.land/bubbles/v2/table"
	"fmt"
)

func (m *Message) Rows() []table.Row {
	var result []table.Row

	for _, i := range m.Items {
		result = append(
			result,
			table.Row{
				i.Label,
				fmt.Sprintf("%.1f", i.Score),
				string(i.Severity),
				i.Detail,
				"",
			},
		)
	}

	return result
}
