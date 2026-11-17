package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"testing"
)

func TestMissingChoiceSetToolQuietFailure(t *testing.T) {
	s := base.New(t)
	message := s.Tools().MustCallToolError(
		constant.DeleteChoiceSet,
		map[string]any{"name": "alfa"},
	)
	assert.StringContains(t, "alfa", message)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestMissingDeviceCustomFieldTargetQuietFailure(t *testing.T) {
	s := base.New(t)
	message := s.Tools().MustCallToolError(
		constant.SetDeviceCustomField,
		map[string]any{
			constant.Device: "alfa",
			constant.Field:  "bravo",
			constant.Values: []any{},
		},
	)
	assert.StringContains(t, "alfa", message)
	assert.Count(t, 0, s.Reporter.Events())
}
