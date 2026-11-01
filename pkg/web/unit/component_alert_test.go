package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component"
	"strings"
	"testing"
)

func TestAnEmptyMessageRaisesNoAlert(t *testing.T) {
	var b strings.Builder
	assert.FatalOnError(t, component.Alert("").Render(&b))
	assert.String(t, "", b.String())
}

func TestAMessageRaisesAnAlert(t *testing.T) {
	var b strings.Builder
	assert.FatalOnError(t, component.Alert("refused").Render(&b))
	assert.StringContains(t, `role="alert"`, b.String())
	assert.StringContains(t, "refused", b.String())
}
