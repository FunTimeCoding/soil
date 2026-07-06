package item

import (
	"charm.land/bubbles/v2/table"
	"github.com/funtimecoding/soil/pkg/bubbletea/style"
)

func New(user bool) *table.Model {
	columns := []table.Column{
		{Title: IdentifierColumn},
		{Title: ScoreColumn},
		{Title: SeverityColumn},
		{Title: DetailColumn},
	}

	if user {
		columns = append(columns, table.Column{Title: UserColumn})
	}

	result := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
	)
	style.Table(&result)

	return &result
}
