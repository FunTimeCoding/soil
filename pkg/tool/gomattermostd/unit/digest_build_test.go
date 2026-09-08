package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"strings"
	"testing"
	"time"
)

func digestAt(hour int, minute int) time.Time {
	return time.Date(2026, 9, 8, hour, minute, 0, 0, time.UTC)
}

func TestDigestBuildEmpty(t *testing.T) {
	assert.String(t, "", digest.Build("papa", nil))
}

func TestDigestBuildSingleMessage(t *testing.T) {
	assert.String(
		t,
		"[papa] 14:32 · Foxtrot: first message",
		digest.Build(
			"papa",
			[]*event.Event{
				event.New(
					constant.MessageEvent,
					"Foxtrot",
					"first message",
					digestAt(14, 32),
				),
			},
		),
	)
}

func TestDigestBuildSingleReaction(t *testing.T) {
	assert.String(
		t,
		"[quebec] 15:12 · Golf reacted ✅",
		digest.Build(
			"quebec",
			[]*event.Event{
				event.New(
					constant.ReactionAddedEvent,
					"Golf",
					"✅",
					digestAt(15, 12),
				),
			},
		),
	)
}

func TestDigestBuildReactionRemoved(t *testing.T) {
	assert.String(
		t,
		"[papa] 15:12 · Golf removed ✅",
		digest.Build(
			"papa",
			[]*event.Event{
				event.New(
					constant.ReactionRemovedEvent,
					"Golf",
					"✅",
					digestAt(15, 12),
				),
			},
		),
	)
}

func TestDigestBuildEnumerated(t *testing.T) {
	assert.Strings(
		t,
		[]string{
			"[papa] 3 new · 14:32-15:10",
			"    Foxtrot: first message",
			"    Foxtrot: second message",
			"    Golf reacted 👀",
		},
		strings.Split(
			digest.Build(
				"papa",
				[]*event.Event{
					event.New(
						constant.MessageEvent,
						"Foxtrot",
						"first message",
						digestAt(14, 32),
					),
					event.New(
						constant.MessageEvent,
						"Foxtrot",
						"second message",
						digestAt(15, 0),
					),
					event.New(
						constant.ReactionAddedEvent,
						"Golf",
						"👀",
						digestAt(15, 10),
					),
				},
			),
			"\n",
		),
	)
}

func TestDigestBuildCollapsesAboveEnumerateLimit(t *testing.T) {
	var events []*event.Event

	for i := range constant.EnumerateLimit + 1 {
		events = append(
			events,
			event.New(
				constant.MessageEvent,
				"Foxtrot",
				"status update",
				digestAt(9, i),
			),
		)
	}

	events = append(
		events,
		event.New(constant.ReactionAddedEvent, "Golf", "✅", digestAt(10, 0)),
	)
	assert.Strings(
		t,
		[]string{
			"[romeo] 6 new · 09:00-10:00 · Foxtrot, Golf · 5 messages, 1 reaction",
			"    latest — Golf reacted ✅",
		},
		strings.Split(digest.Build("romeo", events), "\n"),
	)
}

func TestDigestBuildCollapsesWhenOverBudget(t *testing.T) {
	long := strings.Repeat("a", constant.ExcerptLength)
	var events []*event.Event

	for i := range constant.EnumerateLimit {
		events = append(
			events,
			event.New(constant.MessageEvent, "Foxtrot", long, digestAt(9, i)),
		)
	}

	result := digest.Build("romeo", events)
	assert.True(t, strings.Contains(result, "latest — "))
	assert.Integer(t, 2, len(strings.Split(result, "\n")))
}

func TestDigestBuildTruncatesLongMessage(t *testing.T) {
	result := digest.Build(
		"papa",
		[]*event.Event{
			event.New(
				constant.MessageEvent,
				"Foxtrot",
				strings.Repeat("b", 260),
				digestAt(9, 0),
			),
		},
	)
	assert.True(t, strings.HasSuffix(result, constant.TruncationMarker))
	assert.Integer(t, 200, strings.Count(result, "b"))
}

func TestDigestBuildFlattensNewlines(t *testing.T) {
	assert.String(
		t,
		"[papa] 09:00 · Foxtrot: one two three",
		digest.Build(
			"papa",
			[]*event.Event{
				event.New(
					constant.MessageEvent,
					"Foxtrot",
					"one\n  two\n\nthree",
					digestAt(9, 0),
				),
			},
		),
	)
}

func TestDigestBuildAuthorOverflow(t *testing.T) {
	var events []*event.Event

	for _, name := range []string{"Foxtrot", "Golf", "Hotel", "India", "Juliett"} {
		events = append(
			events,
			event.New(constant.MessageEvent, name, "hello", digestAt(9, 0)),
		)
	}

	assert.True(
		t,
		strings.Contains(digest.Build("romeo", events), "Foxtrot, Golf +3"),
	)
}
