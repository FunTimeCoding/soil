package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/clock"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/web_client"
	"io"
	"testing"
)

func TestWebClientCloseDropsIdleConnections(t *testing.T) {
	s, opened := newCountingServer(t)
	c := web_client.New(clock.New())
	defer c.Close()

	for range 2 {
		r, e := c.Get(s.URL)
		assert.FatalOnError(t, e)
		_, f := io.ReadAll(r.Response.Body)
		assert.FatalOnError(t, f)
		errors.PanicClose(r.Response.Body)
		c.Close()
	}

	assert.Integer(t, int64(2), opened.Load())
}
