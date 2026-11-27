package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/browser/chromium"
	"github.com/funtimecoding/soil/pkg/browser/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/browser/unit/chromium_tester"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"testing"
)

func TestNewClientNeverDialsBrowser(t *testing.T) {
	c := chromium.New("localhost", chromium_tester.ClosedPort(t))
	defer c.Close()
	_, e := c.Targets()
	assert.String(t, "unreachable", classify.Class(e))
}

func TestPageAnswersBrowserThatIsDown(t *testing.T) {
	c := chromium.New("localhost", chromium_tester.ClosedPort(t))
	defer c.Close()
	_, e := protocol.NewIdentifier(c, "alfa").Body()
	assert.String(t, "unreachable", classify.Class(e))
}
