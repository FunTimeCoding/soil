package service

import (
	"fmt"
	"github.com/jackc/pgx/v5"
)

func scanRows(rows pgx.Rows) ([]map[string]any, error) {
	fields := rows.FieldDescriptions()
	var result []map[string]any

	for rows.Next() {
		values, e := rows.Values()

		if e != nil {
			return nil, fmt.Errorf("scan row: %w", e)
		}

		row := make(map[string]any, len(fields))

		for i, f := range fields {
			row[f.Name] = values[i]
		}

		result = append(result, row)
	}

	if e := rows.Err(); e != nil {
		return nil, fmt.Errorf("rows: %w", e)
	}

	return result, nil
}
