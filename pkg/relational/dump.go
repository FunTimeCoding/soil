package relational

import "fmt"

func (d *Database) Dump() map[string]any {
	result := make(map[string]any)

	for _, table := range d.Tables() {
		rows := d.Query(fmt.Sprintf("%s * FROM %s", "SELECT", table))
		rowsResult := make([]map[string]any, 0)

		for rows.Next() {
			columns := rows.FieldDescriptions()
			row := make(map[string]any)

			for i, value := range d.RowValues(rows) {
				row[columns[i].Name] = value
			}

			rowsResult = append(rowsResult, row)
		}

		result[table] = rowsResult
	}

	return result
}
