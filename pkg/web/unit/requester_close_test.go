package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"testing"
)

func TestRequesterCloseDropsIdleConnections(t *testing.T) {
	s, opened := newCountingServer(t)
	r := requester.New(locator.New("unused.example"))
	defer r.Close()
	_, e := r.Bytes(request.Absolute(s.URL))
	assert.FatalOnError(t, e)
	r.Close()
	_, e = r.Bytes(request.Absolute(s.URL))
	assert.FatalOnError(t, e)
	assert.Integer(t, int64(2), opened.Load())
}
