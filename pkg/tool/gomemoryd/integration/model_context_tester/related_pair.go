package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"

func (o *Tester) RelatedPair() {
	o.t.Helper()
	o.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "alpha",
			constant.Content:     "alpha content",
			constant.Description: "alpha description",
		},
	)
	o.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "beta",
			constant.Content:     "beta content",
			constant.Description: "beta description",
		},
	)
}
