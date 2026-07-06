package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (t *Tester) GetTag(path string) string {
	return t.MustCallTool(
		constant.Tag,
		map[string]any{constant.Path: path},
	)
}
