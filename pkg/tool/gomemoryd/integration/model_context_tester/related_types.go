package model_context_tester

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func (o *Tester) RelatedTypes() []string {
	o.t.Helper()
	raw := o.MustCallTool(
		constant.GetMemory,
		map[string]any{constant.MemoryIdentifier: 1},
	)
	var parsed relationTypeResult
	assert.FatalOnError(o.t, json.Unmarshal([]byte(raw), &parsed))
	result := make([]string, 0, len(parsed.Related))

	for _, r := range parsed.Related {
		result = append(result, r.Type)
	}

	return result
}
