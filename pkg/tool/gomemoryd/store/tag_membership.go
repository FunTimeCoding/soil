package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strconv"
	"strings"
)

func (s *Store) TagMembership() (map[string][]int64, error) {
	rows, e := s.database.Query(
		`SELECT t.tag, GROUP_CONCAT(t.memory_identifier)
		FROM memory_tag t
		JOIN memory m ON m.identifier = t.memory_identifier AND m.is_active = 1
		GROUP BY t.tag`,
	)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	result := map[string][]int64{}

	for rows.Next() {
		var tag string
		var joined string

		if f := rows.Scan(&tag, &joined); f != nil {
			return nil, f
		}

		for _, one := range strings.Split(joined, constant.Comma) {
			identifier, f := strconv.ParseInt(one, 10, 64)

			if f == nil {
				result[tag] = append(result[tag], identifier)
			}
		}
	}

	return result, rows.Err()
}
