package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/client_tester"
	"testing"
	"time"
)

func TestTheClientListsUnresolvedThreads(t *testing.T) {
	r := client_tester.New(t, unresolvedUpstream)
	threads, e := r.Client.GetUnresolvedThreads(
		t.Context(),
		"alfa",
		&client.GetUnresolvedThreadsParams{Since: time.UnixMilli(1000)},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, threads)
	assert.String(t, "kilo", threads[0].Root.Identifier)
	assert.String(t, "echo", threads[0].Root.Username)
	assert.Count(t, 1, threads[0].Replies)
	assert.String(t, "lima", threads[0].Replies[0].Identifier)
}

func TestUnresolvedThreadsOfAnUnknownChannelAreNotFound(t *testing.T) {
	r := client_tester.New(t, unresolvedUpstream)
	_, e := r.Client.GetUnresolvedThreads(
		t.Context(),
		"missing",
		&client.GetUnresolvedThreadsParams{Since: time.UnixMilli(1000)},
	)
	assert.True(t, not_found.Is(e))
}
