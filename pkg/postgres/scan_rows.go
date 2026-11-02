package postgres

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
	"github.com/jackc/pgx/v5"
)

func scanRows(rows pgx.Rows) (*result.Result, error) {
	var columns []string

	for _, f := range rows.FieldDescriptions() {
		columns = append(columns, f.Name)
	}

	var values [][]any

	for rows.Next() {
		v, e := rows.Values()

		if e != nil {
			return nil, fmt.Errorf("scan row: %w", e)
		}

		values = append(values, v)
	}

	if e := rows.Err(); e != nil {
		return nil, fmt.Errorf("rows: %w", e)
	}

	return result.New(columns, values), nil
}
