package base

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/model_context/response"
)

func (s *Server) RelatedTypes() []string {
	s.t.Helper()
	raw := s.Tools().MustCallTool(
		constant.GetMemory,
		map[string]any{constant.MemoryIdentifier: 1},
	)
	var parsed response.SlimMemoryWithHistory
	assert.FatalOnError(s.t, json.Unmarshal([]byte(raw), &parsed))
	result := make([]string, 0, len(parsed.Related))

	for _, r := range parsed.Related {
		result = append(result, r.Type)
	}

	return result
}
