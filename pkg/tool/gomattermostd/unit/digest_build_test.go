package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/digest_tester"
	"strings"
	"testing"
)

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
					digest_tester.At(14, 32),
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
					digest_tester.At(15, 12),
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
					digest_tester.At(15, 12),
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
						digest_tester.At(14, 32),
					),
					event.New(
						constant.MessageEvent,
						"Foxtrot",
						"second message",
						digest_tester.At(15, 0),
					),
					event.New(
						constant.ReactionAddedEvent,
						"Golf",
						"👀",
						digest_tester.At(15, 10),
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
				digest_tester.At(9, i),
			),
		)
	}

	events = append(
		events,
		event.New(
			constant.ReactionAddedEvent,
			"Golf",
			"✅",
			digest_tester.At(10, 0),
		),
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

func TestDigestBuildCollapsesWhenOverLimit(t *testing.T) {
	long := strings.Repeat("a", constant.ExcerptLength)
	var events []*event.Event

	for i := range constant.EnumerateLimit {
		events = append(
			events,
			event.New(
				constant.MessageEvent,
				"Foxtrot",
				long,
				digest_tester.At(9, i),
			),
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
				digest_tester.At(9, 0),
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
					digest_tester.At(9, 0),
				),
			},
		),
	)
}

func TestDigestBuildAuthorOverflow(t *testing.T) {
	var events []*event.Event

	for _, name := range []string{
		"Foxtrot",
		"Golf",
		"Hotel",
		"India",
		"Juliett",
	} {
		events = append(
			events,
			event.New(
				constant.MessageEvent,
				name,
				"hello",
				digest_tester.At(9, 0),
			),
		)
	}

	assert.True(
		t,
		strings.Contains(digest.Build("romeo", events), "Foxtrot, Golf +3"),
	)
}
