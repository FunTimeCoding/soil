package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (t *Tester) SetTag(
	path string,
	sourceType string,
) string {
	return t.MustCallTool(
		constant.Tag,
		map[string]any{
			constant.Path:       path,
			constant.SourceType: sourceType,
		},
	)
}
