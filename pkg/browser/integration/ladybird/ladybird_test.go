//go:build browser

package ladybird

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/browser/ladybird"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/ladybird_option"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"testing"
	"time"
)

func TestLadybirdDrivesPage(t *testing.T) {
	events := make(chan *message.Reply, 64)
	o := ladybird_option.New()
	o.Event = func(r *message.Reply) { events <- r }
	c := ladybird.New(o)
	c.Open()
	t.Cleanup(c.Close)
	contexts, e := c.Tree()
	assert.FatalOnError(t, e)
	context := contexts[0].Identifier
	assert.FatalOnError(t, c.Subscribe("log.entryAdded"))
	assert.FatalOnError(
		t,
		c.Navigate(
			context,
			"data:text/html,<title>probe</title><h1>hello</h1>",
		),
	)
	result, e := c.Evaluate(context, "console.log('from test'); document.title")
	assert.FatalOnError(t, e)
	assert.String(t, `{"type":"string","value":"probe"}`, string(result))

	select {
	case r := <-events:
		assert.String(t, "log.entryAdded", r.Method)
		assert.StringContains(t, "from test", string(r.Parameters))
	case <-time.After(5 * time.Second):
		t.Fatal("no console event")
	}

	image, e := c.Screenshot(context)
	assert.FatalOnError(t, e)
	assert.String(t, "\x89PNG", string(image[:4]))
	_, e = c.Evaluate(context, "missing.property")
	assert.True(t, validation.Is(e))
}
