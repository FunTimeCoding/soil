package publish_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Publish() string {
	o.t.Helper()
	commit, _, e := o.Server.Service.Publish()
	assert.FatalOnError(o.t, e)

	return commit
}
