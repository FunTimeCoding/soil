package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (t *Tester) PushWithMetadata(
	collection string,
	path string,
	body string,
	metadata map[string]string,
) string {
	arguments := map[string]any{
		constant.Collection: collection,
		constant.Path:       path,
		constant.Body:       body,
	}

	if len(metadata) > 0 {
		arguments[constant.Metadata] = metadata
	}

	return t.MustCallTool(constant.Push, arguments)
}
