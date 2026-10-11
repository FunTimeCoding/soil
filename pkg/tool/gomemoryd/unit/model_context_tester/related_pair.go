package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"

func (o *Tester) RelatedPair() {
	o.t.Helper()
	o.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "alfa",
			constant.Content:     "alfa content",
			constant.Description: "alfa description",
		},
	)
	o.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "bravo",
			constant.Content:     "bravo content",
			constant.Description: "bravo description",
		},
	)
}
