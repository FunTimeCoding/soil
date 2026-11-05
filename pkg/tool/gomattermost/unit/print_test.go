package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/terminal/unit/command_end_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gomattermost"
	"github.com/funtimecoding/soil/pkg/tool/gomattermost/thread_summary"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	"strings"
	"testing"
	"time"
)

func TestPrintShowsTheWindowEachThreadAndItsSummary(t *testing.T) {
	var output bytes.Buffer
	var failure bytes.Buffer
	gomattermost.Print(
		terminal.NewWith(command_end_recorder.New(), &output, &failure, nil),
		time.Date(2026, 11, 4, 0, 0, 0, 0, time.UTC),
		[]*thread_summary.Summary{
			thread_summary.New(
				client.Thread{
					Root: client.Post{Username: "lima", Message: "alfa is down"},
				},
				"kilo",
			),
			thread_summary.New(
				client.Thread{
					Root: client.Post{
						Username: "mike",
						Message:  "bravo is slow",
					},
				},
				"papa",
			),
		},
	)
	text := output.String()
	assert.String(t, "", failure.String())
	assert.True(
		t,
		strings.HasPrefix(text, "Unresolved threads since 2026-11-04 00:00\n"),
	)
	alfa := strings.Index(text, "alfa is down")
	kilo := strings.Index(text, "kilo")
	bravo := strings.Index(text, "bravo is slow")
	papa := strings.Index(text, "papa")
	assert.True(t, alfa > 0 && alfa < kilo)
	assert.True(t, kilo < bravo && bravo < papa)
	assert.StringNotContains(t, "No unresolved threads", text)
}

func TestPrintSaysWhenNothingIsUnresolved(t *testing.T) {
	var output bytes.Buffer
	gomattermost.Print(
		terminal.NewWith(command_end_recorder.New(), &output, &output, nil),
		time.Date(2026, 11, 4, 0, 0, 0, 0, time.UTC),
		nil,
	)
	assert.String(
		t,
		"Unresolved threads since 2026-11-04 00:00\nNo unresolved threads\n",
		output.String(),
	)
}
