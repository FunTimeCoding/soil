package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (t *Tester) RemoveContext(
	collection string,
	pathPrefix string,
) string {
	return t.MustCallTool(
		constant.RemoveContext,
		map[string]any{
			constant.Collection: collection,
			constant.PathPrefix: pathPrefix,
		},
	)
}
