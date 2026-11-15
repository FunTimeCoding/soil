package model_context_tester

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func Decode(
	t *testing.T,
	raw string,
) map[string]any {
	t.Helper()
	var result map[string]any
	assert.FatalOnError(t, json.Unmarshal([]byte(raw), &result))

	return result
}
