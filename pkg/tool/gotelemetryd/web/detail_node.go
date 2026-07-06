package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"maragu.dev/gomponents"
	"sort"
)

func detailNode(detail map[string]string) gomponents.Node {
	keys := make([]string, 0, len(detail))

	for k := range detail {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var parts []string

	for _, k := range keys {
		parts = append(parts, key_value.Equals(k, detail[k]))
	}

	return gomponents.Text(join.CommaSpace(parts))
}
