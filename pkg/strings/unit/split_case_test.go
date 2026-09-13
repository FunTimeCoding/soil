package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestSplitCase(t *testing.T) {
	assert.String(t, "noise With Match", strings.SplitCase("noiseWithMatch"))
	assert.String(t, "HTTP Server", strings.SplitCase("HTTPServer"))
	assert.String(t, "plain", strings.SplitCase("plain"))
	assert.String(t, "already split", strings.SplitCase("already split"))
	assert.String(t, "value2 Next", strings.SplitCase("value2Next"))
	assert.String(t, "", strings.SplitCase(""))
}
