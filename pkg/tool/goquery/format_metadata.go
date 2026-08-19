package goquery

import (
	"fmt"
	"sort"
	"strings"
)

func formatMetadata(metadata map[string][]string) string {
	keys := make([]string, 0, len(metadata))

	for k := range metadata {
		if k == "source_type" {
			continue
		}

		keys = append(keys, k)
	}

	sort.Strings(keys)
	var parts []string

	for _, k := range keys {
		parts = append(
			parts,
			fmt.Sprintf("%s=%s", k, strings.Join(metadata[k], ",")),
		)
	}

	return strings.Join(parts, "  ")
}
