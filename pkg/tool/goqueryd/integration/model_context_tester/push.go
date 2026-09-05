package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (t *Tester) Push(
	collection string,
	path string,
	body string,
) string {
	return t.MustCallTool(
		constant.Push,
		map[string]any{
			constant.Collection: collection,
			constant.Path:       path,
			constant.Body:       body,
		},
	)
}
