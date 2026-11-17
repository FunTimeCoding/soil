package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Subscribe(
	root string,
	callsign string,
	alias string,
) {
	o.t.Helper()
	p, e := o.Mattermost.FindPost(root)
	assert.FatalOnError(o.t, e)
	_, _, f := o.Service().Subscribe(p, callsign, alias, "")
	assert.FatalOnError(o.t, f)
}
