package publish_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func (o *Tester) Publish(t *testing.T) string {
	t.Helper()
	commit, _, e := o.Server.Service.Publish()
	assert.FatalOnError(t, e)

	return commit
}
