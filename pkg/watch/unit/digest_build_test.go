package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/watch/constant"
	"github.com/funtimecoding/soil/pkg/watch/digest"
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
	"github.com/funtimecoding/soil/pkg/watch/unit/digest_tester"
	"strings"
	"testing"
)

func TestDigestBuildEmpty(t *testing.T) {
	assert.String(t, "", digest.Build("papa", nil))
}

func TestDigestBuildSingle(t *testing.T) {
	assert.String(
		t,
		"[papa] 14:32 · Foxtrot: first message",
		digest.Build(
			"papa",
			[]*event.Event{
				event.New(
					constant.MessageNoun,
					"Foxtrot",
					"Foxtrot: first message",
					digest_tester.At(14, 32),
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
			"    Golf reacted :eyes:",
		},
		strings.Split(
			digest.Build(
				"papa",
				[]*event.Event{
					event.New(
						constant.MessageNoun,
						"Foxtrot",
						"Foxtrot: first message",
						digest_tester.At(14, 32),
					),
					event.New(
						constant.MessageNoun,
						"Foxtrot",
						"Foxtrot: second message",
						digest_tester.At(15, 0),
					),
					event.New(
						constant.ReactionNoun,
						"Golf",
						"Golf reacted :eyes:",
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
				constant.MessageNoun,
				"Foxtrot",
				"Foxtrot: status update",
				digest_tester.At(9, i),
			),
		)
	}

	events = append(
		events,
		event.New(
			constant.ReactionNoun,
			"Golf",
			"Golf reacted :white_check_mark:",
			digest_tester.At(10, 0),
		),
	)
	assert.Strings(
		t,
		[]string{
			"[romeo] 6 new · 09:00-10:00 · Foxtrot, Golf · 5 messages, 1 reaction",
			"    latest — Golf reacted :white_check_mark:",
		},
		strings.Split(digest.Build("romeo", events), "\n"),
	)
}

func TestDigestCountsInFirstSeenOrder(t *testing.T) {
	var events []*event.Event

	for i, noun := range []string{
		"comment",
		"approval",
		"comment",
		"approval",
		"comment",
	} {
		events = append(
			events,
			event.New(
				noun,
				"Hotel",
				"Hotel did something",
				digest_tester.At(9, i),
			),
		)
	}

	assert.True(
		t,
		strings.Contains(
			digest.Build("sierra", events),
			"3 comments, 2 approvals",
		),
	)
}

func TestDigestCountsPluralizeSibilantNouns(t *testing.T) {
	var events []*event.Event

	for i := range 6 {
		events = append(
			events,
			event.New("push", "Hotel", "pushed", digest_tester.At(9, i)),
		)
	}

	assert.True(t, strings.Contains(digest.Build("sierra", events), "6 pushes"))
}

func TestDigestAuthorsSkipAnonymousEvents(t *testing.T) {
	var events []*event.Event

	for i, author := range []string{"", "Hotel", "", "India"} {
		events = append(
			events,
			event.New(
				"pipeline",
				author,
				"pipeline 11 failed",
				digest_tester.At(9, i),
			),
		)
	}

	events = append(
		events,
		event.New("pipeline", "", "pipeline 12 running", digest_tester.At(9, 5)),
	)
	events = append(
		events,
		event.New("pipeline", "", "pipeline 12 failed", digest_tester.At(9, 6)),
	)
	assert.String(
		t,
		"[sierra] 6 new · 09:00-09:06 · Hotel, India · 6 pipelines",
		strings.Split(digest.Build("sierra", events), "\n")[0],
	)
}

func TestDigestBuildCollapsesWhenOverLimit(t *testing.T) {
	long := strings.Repeat("a", constant.ExcerptLength)
	var events []*event.Event

	for i := range constant.EnumerateLimit {
		events = append(
			events,
			event.New(
				constant.MessageNoun,
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
				constant.MessageNoun,
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

func TestExcerptTruncatesLongText(t *testing.T) {
	result := digest.Excerpt(strings.Repeat("b", 260))
	assert.True(t, strings.HasSuffix(result, constant.TruncationMarker))
	assert.Integer(t, 200, strings.Count(result, "b"))
}

func TestExcerptFlattensNewlines(t *testing.T) {
	assert.String(t, "one two three", digest.Excerpt("one\n  two\n\nthree"))
}
