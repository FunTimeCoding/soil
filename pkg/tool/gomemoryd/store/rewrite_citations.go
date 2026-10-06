package store

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"strings"
)

func rewriteCitations(
	t *sql.Tx,
	from string,
	to string,
	source string,
	now string,
) ([]int64, error) {
	rows, e := t.Query(
		`SELECT identifier, name, content, description FROM memory
		WHERE is_active = 1 AND provenance_file = '' AND instr(content, ?) > 0`,
		from,
	)

	if e != nil {
		return nil, e
	}

	var citing []*record.Version

	for rows.Next() {
		v := record.NewVersion()

		if f := rows.Scan(
			&v.MemoryIdentifier,
			&v.Name,
			&v.Content,
			&v.Description,
		); f != nil {
			errors.LogClose(rows)

			return nil, f
		}

		citing = append(citing, v)
	}

	errors.LogClose(rows)
	var result []int64

	for _, v := range citing {
		v.Content = strings.ReplaceAll(v.Content, from, to)

		if f := rewriteCitation(t, v, source, now); f != nil {
			return nil, f
		}

		result = append(result, v.MemoryIdentifier)
	}

	return result, nil
}
