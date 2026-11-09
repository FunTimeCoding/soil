package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/model_context_tester"
	"testing"
)

func TestKeyToUnknownSessionIsAnsweredQuietly(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	assert.String(
		t,
		"step 1: session alfa not found",
		o.Client.MustCallToolError(
			constant.SendKey,
			map[string]any{"session_id": "alfa", "keys": []string{"Return"}},
		),
	)
}
