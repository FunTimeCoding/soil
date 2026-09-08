package mattermost_client_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (t *Tester) Drop() {
	t.t.Helper()
	t.mutex.Lock()
	defer t.mutex.Unlock()
	assert.FatalOnError(t.t, t.connection.Close())
}
