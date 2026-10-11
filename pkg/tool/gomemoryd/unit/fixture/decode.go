package fixture

import (
	"encoding/json"
	"testing"

	"github.com/funtimecoding/soil/pkg/assert"
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
