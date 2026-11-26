package service_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
)

func (o *Tester) Page(
	title string,
	buttons ...string,
) {
	o.t.Helper()
	items := []any{}
	values := []any{nil}

	for i, name := range buttons {
		items = append(
			items,
			map[string]any{"i": i, "r": "button", "n": name, "c": []any{}},
		)
		values = append(
			values,
			map[string]any{"type": "node", "sharedId": fmt.Sprintf("ref-%d", i)},
		)
	}

	items = append(
		items,
		map[string]any{"r": "text", "n": "idle", "c": []any{}},
	)
	values[0] = map[string]any{
		"type": "string",
		"value": notation.Encode(
			map[string]any{"title": title, "tree": items},
			false,
		),
	}
	o.Client.SetResult(
		constant.SnapshotScript,
		notation.Encode(map[string]any{"type": "array", "value": values}, false),
	)
}
