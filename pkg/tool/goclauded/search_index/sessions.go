package search_index

import "github.com/funtimecoding/soil/pkg/errors"

func (x *Index) Sessions() ([]string, error) {
	rows, e := x.database.Query("SELECT session FROM harbor_file")

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(rows)
	var result []string

	for rows.Next() {
		var session string

		if f := rows.Scan(&session); f != nil {
			return nil, f
		}

		result = append(result, session)
	}

	return result, rows.Err()
}
