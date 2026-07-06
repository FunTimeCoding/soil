package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (t *Tester) AddContext(
	collection string,
	pathPrefix string,
	description string,
) string {
	return t.MustCallTool(
		constant.AddContext,
		map[string]any{
			constant.Collection:  collection,
			constant.PathPrefix:  pathPrefix,
			constant.Description: description,
		},
	)
}
