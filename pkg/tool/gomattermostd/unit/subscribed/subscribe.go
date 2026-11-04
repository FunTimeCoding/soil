package subscribed

import (
	"github.com/funtimecoding/soil/pkg/assert"
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"testing"
)

func (o *Subscribed) Subscribe(
	t *testing.T,
	root string,
	callsign string,
	d goclaudedConstant.Delivery,
) {
	t.Helper()
	o.Mattermost.AddPost(root)
	o.Mattermost.AddUser("foxtrot", "Foxtrot")
	p, e := o.Mattermost.FindPost(root)
	assert.FatalOnError(t, e)
	_, _, f := o.Service.Subscribe(p, callsign, "papa", d)
	assert.FatalOnError(t, f)
}
