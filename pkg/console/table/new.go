package table

import "github.com/funtimecoding/soil/pkg/console/types/table_column"

func New(headers ...string) *Table {
	result := &Table{}

	for _, h := range headers {
		result.columns = append(result.columns, table_column.New(h, len(h)))
	}

	return result
}
