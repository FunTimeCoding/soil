package item

import (
	"charm.land/bubbles/v2/table"
	"github.com/funtimecoding/soil/pkg/bubbletea/constant"
	"github.com/funtimecoding/soil/pkg/bubbletea/style"
)

func New(user bool) *table.Model {
	columns := []table.Column{
		{Title: constant.ItemIdentifierColumn},
		{Title: constant.ItemScoreColumn},
		{Title: constant.ItemSeverityColumn},
		{Title: constant.ItemDetailColumn},
	}

	if user {
		columns = append(
			columns,
			table.Column{Title: constant.ItemUserColumn},
		)
	}

	result := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
	)
	style.Table(&result)

	return &result
}
