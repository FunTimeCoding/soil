package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (t *Tester) ListDocuments(collection string) string {
	return t.MustCallTool(
		constant.List,
		map[string]any{constant.Collection: collection},
	)
}
