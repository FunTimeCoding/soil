package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"sort"
)

func metadataJoins(metadata map[string]string) (string, []any) {
	keys := make([]string, 0, len(metadata))

	for key := range metadata {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	var parts []string
	var arguments []any

	for i, key := range keys {
		alias := fmt.Sprintf("m%d", i)
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
		arguments = append(arguments, key, metadata[key])
	}

	return join.Space(parts...), arguments
}
