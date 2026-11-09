package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"testing"
)

func TestNewClientNeverDialsBrowser(t *testing.T) {
	c := chromium.New("localhost", closedPort(t))
	defer c.Close()
	_, e := c.Targets()
	assert.String(t, "unreachable", classify.Class(e))
}

func TestPageAnswersBrowserThatIsDown(t *testing.T) {
	c := chromium.New("localhost", closedPort(t))
	defer c.Close()
	_, e := protocol.NewIdentifier(c, "alfa").Body()
	assert.String(t, "unreachable", classify.Class(e))
}
