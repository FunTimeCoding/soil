package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestLabelReservedKeyBlocked(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "labeling")

	for _, key := range constant.ReservedLabelKeys {
		result := a.MustCallToolError(
			constant.Label,
			map[string]any{
				constant.Key:   key,
				constant.Value: "should be blocked",
			},
		)
		assert.StringContains(t, "reserved key", result)
		assert.StringContains(t, constant.EditSession, result)
	}
}

func TestLabelCustomKey(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "labeling")
	a.MustCallTool(
		constant.Label,
		map[string]any{
			constant.Key:   "stage",
			constant.Value: "canary",
		},
	)
	result := a.MustCallTool(
		constant.Label,
		map[string]any{constant.Key: "stage"},
	)
	assert.StringContains(t, "stage=canary", result)
}
