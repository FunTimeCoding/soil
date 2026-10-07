package table

import "github.com/funtimecoding/soil/pkg/console/types/table_column"

type Table struct {
	columns []*table_column.Column
	rows    [][]string
}
