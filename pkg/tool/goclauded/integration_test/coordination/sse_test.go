package coordination

import (
	"bufio"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestSSEPushesRosterOnConnect(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	events := connectSSE(t, s)
	e := nextEvent(t, events)
	assert.String(t, "roster", e.name)
	assert.True(t, e.payload != "")
}

func TestSSEPushesActivityOnConnect(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	events := connectSSE(t, s)
	nextEvent(t, events)
	e := nextEvent(t, events)
	assert.String(t, "activity", e.name)
}

func TestSSEPushesAfterMutation(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	events := connectSSE(t, s)
	nextEvent(t, events)
	nextEvent(t, events)
	a.Announce(a.Name(), "trigger notification")
	e := nextEvent(t, events)
	assert.String(t, "roster", e.name)
	assert.StringContains(t, "trigger notification", e.payload)
}

func TestSSEMultiLineBody(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "setup")
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "line one\nline two\nline three",
		},
	)
	events := connectSSE(t, s)
	nextEvent(t, events)
	activity := nextEvent(t, events)
	assert.String(t, "activity", activity.name)
	assert.StringContains(t, "line one", activity.payload)
	assert.StringContains(t, "line two", activity.payload)
}

type sseEvent struct {
	name    string
	payload string
}

func connectSSE(
	t *testing.T,
	s *base.Server,
) <-chan sseEvent {
	t.Helper()
	l := fmt.Sprintf(
		"http://localhost:%d/event?subscribe=roster,activity,summary",
		s.Port(),
	)
	r, e := http.Get(l)
	assert.FatalOnError(t, e)
	t.Cleanup(func() { errors.PanicClose(r.Body) })
	events := make(chan sseEvent, 10)
	go func() {
		scanner := bufio.NewScanner(r.Body)
		var name string
		var lines []string

		for scanner.Scan() {
			line := scanner.Text()

			if strings.HasPrefix(line, "event: ") {
				name = strings.TrimPrefix(line, "event: ")
				lines = nil
			} else if strings.HasPrefix(line, "data: ") {
				lines = append(
					lines,
					strings.TrimPrefix(line, "data: "),
				)
			} else if line == "" && name != "" {
				events <- sseEvent{
					name:    name,
					payload: join.NewLine(lines),
				}
				name = ""
				lines = nil
			}
		}

		close(events)
	}()

	return events
}

func nextEvent(
	t *testing.T,
	events <-chan sseEvent,
) sseEvent {
	t.Helper()

	select {
	case e := <-events:
		return e
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for SSE event")

		return sseEvent{}
	}
}
