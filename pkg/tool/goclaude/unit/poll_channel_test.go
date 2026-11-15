package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/unit/channel_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestPollChannelStallsOnceAtThreshold(t *testing.T) {
	u := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusInternalServerError),
	)
	c := client_tester.New(t, u.URL)
	s, k := channel_tester.New()
	failures := 0

	for range 5 {
		failures = goclaude.PollChannel(c, s, "Ash", failures)
	}

	assert.Integer(t, 5, failures)
	stalled := k.Kind("stalled")
	assert.Count(t, 1, stalled)
	assert.String(
		t,
		"channel delivery is stalled: 3 consecutive poll failures reaching the coordination daemon. Nothing is being delivered here, and pending traffic is waiting for your next prompt instead.",
		stalled[0].Content,
	)
}

func TestPollChannelResumesAndDeliversAfterStall(t *testing.T) {
	u := upstream_server.New(
		t,
		upstream_server.WithStatus(
			http.StatusInternalServerError,
			http.StatusInternalServerError,
			http.StatusInternalServerError,
			http.StatusOK,
		),
		upstream_server.WithAnswer(constant.FixtureChannelEntry),
		upstream_server.WithObject(),
	)
	c := client_tester.New(t, u.URL)
	s, k := channel_tester.New()
	failures := 0

	for range 4 {
		failures = goclaude.PollChannel(c, s, "Ash", failures)
	}

	assert.Integer(t, 0, failures)
	stalled := k.Kind("stalled")
	assert.Count(t, 2, stalled)
	assert.String(t, "channel delivery resumed.", stalled[1].Content)
	messages := k.Kind("message")
	assert.Count(t, 1, messages)
	assert.String(t, "Ash: hello", messages[0].Content)
}

func TestPollChannelSuccessResetsCount(t *testing.T) {
	u := upstream_server.New(
		t,
		upstream_server.WithStatus(
			http.StatusInternalServerError,
			http.StatusInternalServerError,
			http.StatusOK,
			http.StatusInternalServerError,
			http.StatusInternalServerError,
		),
		upstream_server.WithAnswer(constant.FixtureChannelEntry),
		upstream_server.WithObject(),
	)
	c := client_tester.New(t, u.URL)
	s, k := channel_tester.New()
	failures := 0

	for range 5 {
		failures = goclaude.PollChannel(c, s, "Ash", failures)
	}

	assert.Integer(t, 2, failures)
	assert.Count(t, 0, k.Kind("stalled"))
}
