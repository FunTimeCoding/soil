package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestRequesterCloseDropsIdleConnections(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer("counted"))
	r := requester.New(locator.New("unused.example"))
	defer r.Close()
	_, e := r.Bytes(request.Absolute(s.URL))
	assert.FatalOnError(t, e)
	r.Close()
	_, e = r.Bytes(request.Absolute(s.URL))
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, s.Opened())
}
