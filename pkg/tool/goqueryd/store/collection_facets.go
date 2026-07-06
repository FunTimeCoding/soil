package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"sort"
)

func (s *Store) CollectionFacets(
	collection string,
	metadata map[string]string,
	threshold int,
	keys ...string,
) []Facet {
	var parts []string
	var arguments []any
	parts = append(
		parts,
		`SELECT m.key, m.value, COUNT(*)
		FROM document_metadata m
		JOIN document d
		ON d.identifier = m.document_identifier`,
	)
	i := 0

	for key, value := range metadata {
		alias := fmt.Sprintf("mf%d", i)
		parts = append(
			parts,
			fmt.Sprintf(
				`JOIN document_metadata %s
				ON %s.document_identifier = d.identifier
				AND %s.key = ? AND %s.value = ?`,
				alias,
				alias,
				alias,
				alias,
			),
		)
		arguments = append(arguments, key, value)
		i++
	}

	parts = append(parts, `WHERE d.collection = ? AND d.active = 1`)
	arguments = append(arguments, collection)

	if len(keys) > 0 {
		parts = append(
			parts,
			fmt.Sprintf(`AND m.key IN (%s)`, placeholders(len(keys))),
		)

		for _, k := range keys {
			arguments = append(arguments, k)
		}
	}

	parts = append(
		parts,
		`GROUP BY m.key, m.value
		ORDER BY m.key, count(*) DESC`,
	)
	rows, e := s.database.Query(join.Space(parts...), arguments...)
	errors.PanicOnError(e)
	defer errors.PanicClose(rows)
	counts := map[string]map[string]int{}

	for rows.Next() {
		var key, value string
		var count int
		errors.PanicOnError(rows.Scan(&key, &value, &count))

		if counts[key] == nil {
			counts[key] = map[string]int{}
		}

		counts[key][value] = count
	}

	errors.PanicOnError(rows.Err())
	sorted := make([]string, 0, len(counts))

	for k := range counts {
		sorted = append(sorted, k)
	}

	sort.Strings(sorted)
	var result []Facet

	for _, key := range sorted {
		values := counts[key]
		f := Facet{
			Key:      key,
			Distinct: len(values),
		}

		if f.Distinct <= threshold {
			f.Values = values
		}

		result = append(result, f)
	}

	return result
}
