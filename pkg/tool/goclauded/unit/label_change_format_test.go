package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/label_change"
	"testing"
)

func TestLabelChangeFormatFromUnset(t *testing.T) {
	assert.String(
		t,
		"colour (unset)→green",
		label_change.Format("colour", "", "green"),
	)
}

func TestLabelChangeFormatReplace(t *testing.T) {
	assert.String(
		t,
		"colour blue→green",
		label_change.Format("colour", "blue", "green"),
	)
}

func TestLabelChangeFormatRemove(t *testing.T) {
	assert.String(
		t,
		"colour blue→ (unset)",
		label_change.Format("colour", "blue", ""),
	)
}
