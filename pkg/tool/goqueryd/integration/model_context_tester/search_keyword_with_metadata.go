package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (t *Tester) SearchKeywordWithMetadata(
	query string,
	metadata map[string]string,
) string {
	return t.MustCallTool(
		constant.Search,
		map[string]any{
			"query":           query,
			constant.Mode:     "keyword",
			constant.Metadata: metadata,
		},
	)
}
