package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (t *Tester) SearchKeyword(query string) string {
	return t.MustCallTool(
		constant.Search,
		map[string]any{
			"query": query,
			"mode":  "keyword",
		},
	)
}
