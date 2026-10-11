package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer_tester"
	"strings"
	"testing"
)

func TestPointersHeadingLinkLive(t *testing.T) {
	l := pointer_tester.Headings(
		map[string]string{"doc/guide/other.md": "# Other\n\n## Part\n"},
		"doc/guide/other.md",
	)(
		"doc/guide/index.md",
		strings.NewReader("See [other](other.md#part) there.\n"),
	)
	assertReport(t, "doc/guide/index.md", false, nil, "", l)
}

func TestPointersHeadingLinkDead(t *testing.T) {
	l := pointer_tester.Headings(
		map[string]string{"doc/guide/other.md": "# Other\n\n## Part\n"},
		"doc/guide/other.md",
	)(
		"doc/guide/index.md",
		strings.NewReader("See [other](other.md#parts) there.\n"),
	)
	assertReport(
		t,
		"doc/guide/index.md",
		true,
		pointer_tester.DeadHeading(
			"doc/guide/index.md",
			1,
			"See [other](other.md#parts) there.",
			"Referenced heading does not exist - nearest: #part (\"Part\"), #other (\"Other\")",
		),
		"",
		l,
	)
}

func TestPointersHeadingSpanLive(t *testing.T) {
	l := pointer_tester.Headings(
		map[string]string{"doc/guide/other.md": "# Other\n\n## Part\n"},
		"doc/guide/other.md",
	)(
		"doc/index.md",
		strings.NewReader("Read `doc/guide/other.md#part` first.\n"),
	)
	assertReport(t, "doc/index.md", false, nil, "", l)
}

func TestPointersHeadingSpanDead(t *testing.T) {
	l := pointer_tester.Headings(
		map[string]string{"doc/guide/other.md": "## Part\n"},
		"doc/guide/other.md",
	)(
		"doc/index.md",
		strings.NewReader("Read `doc/guide/other.md#gone` first.\n"),
	)
	assertReport(
		t,
		"doc/index.md",
		true,
		pointer_tester.DeadHeading(
			"doc/index.md",
			1,
			"Read `doc/guide/other.md#gone` first.",
			"Referenced heading does not exist - nearest: #part (\"Part\")",
		),
		"",
		l,
	)
}

func TestPointersHeadingSameFileLive(t *testing.T) {
	content := "## Top\n\nBack to [top](#top).\n"
	l := pointer_tester.Headings(
		map[string]string{"doc/guide/index.md": content},
		"doc/guide/index.md",
	)("doc/guide/index.md", strings.NewReader(content))
	assertReport(t, "doc/guide/index.md", false, nil, "", l)
}

func TestPointersHeadingSameFileDead(t *testing.T) {
	content := "## Top\n\nBack to [top](#bottom).\n"
	l := pointer_tester.Headings(
		map[string]string{"doc/guide/index.md": content},
		"doc/guide/index.md",
	)("doc/guide/index.md", strings.NewReader(content))
	assertReport(
		t,
		"doc/guide/index.md",
		true,
		pointer_tester.DeadHeading(
			"doc/guide/index.md",
			3,
			"Back to [top](#bottom).",
			"Referenced heading does not exist - nearest: #top (\"Top\")",
		),
		"",
		l,
	)
}

func TestPointersHeadingSpanOpeningWithHashIsNoAnchor(t *testing.T) {
	l := pointer_tester.Headings(
		map[string]string{"doc/index.md": "# Index\n"},
		"doc/index.md",
	)(
		"doc/index.md",
		strings.NewReader("Open `#settings/usage` in the browser.\n"),
	)
	assertReport(t, "doc/index.md", false, nil, "", l)
}

func TestPointersHeadingLinkInsideCodeSpanIsNoLink(t *testing.T) {
	l := pointer_tester.Headings(
		map[string]string{"doc/index.md": "# Index\n"},
		"doc/index.md",
	)(
		"doc/index.md",
		strings.NewReader("Write anchors as `[top](#top)` links.\n"),
	)
	assertReport(t, "doc/index.md", false, nil, "", l)
}

func TestPointersHeadingOnNonMarkdownTarget(t *testing.T) {
	l := pointer_tester.Headings(nil, "pkg/lint/check.go")(
		"doc/index.md",
		strings.NewReader("See `pkg/lint/check.go#pointers` for the loop.\n"),
	)
	assertReport(
		t,
		"doc/index.md",
		true,
		pointer_tester.FragmentTarget(
			"doc/index.md",
			"See `pkg/lint/check.go#pointers` for the loop.",
		),
		"",
		l,
	)
}

func TestPointersHeadingOnDirectory(t *testing.T) {
	l := pointer_tester.Headings(nil, "doc/guide")(
		"doc/index.md",
		strings.NewReader("Everything in `doc/guide/#part` applies.\n"),
	)
	assertReport(
		t,
		"doc/index.md",
		true,
		pointer_tester.FragmentTarget(
			"doc/index.md",
			"Everything in `doc/guide/#part` applies.",
		),
		"",
		l,
	)
}

func TestPointersHeadingOnDeadFileReportsThePath(t *testing.T) {
	l := pointer_tester.Headings(nil)(
		"doc/index.md",
		strings.NewReader("Read `doc/gone.md#part` first.\n"),
	)
	assertReport(
		t,
		"doc/index.md",
		true,
		pointer_tester.Dead("doc/index.md", "Read `doc/gone.md#part` first.", 1),
		"",
		l,
	)
}

func TestPointersHeadingUnreadableIsTallied(t *testing.T) {
	l, seen := pointer_tester.Recording("doc/guide/other.md")
	assertReport(
		t,
		"doc/index.md",
		false,
		nil,
		"",
		l(
			"doc/index.md",
			strings.NewReader("Read `doc/guide/other.md#part` first.\n"),
		),
	)
	assert.Any(
		t,
		[]string{"doc/index.md:1 heading doc/guide/other.md#part"},
		*seen,
	)
}

func TestPointersHeadingOnLocatorIsLeftAlone(t *testing.T) {
	l := pointer_tester.Headings(nil)(
		"doc/index.md",
		strings.NewReader(
			"---\nhosts: alfa.example\n---\nSee [the page](https://alfa.example/page#part).\n",
		),
	)
	assertReport(t, "doc/index.md", false, nil, "", l)
}

func TestPointersHeadingEncodedLinkLive(t *testing.T) {
	l := pointer_tester.Headings(
		map[string]string{"doc/guide/two words.md": "# Two\n\n## Part\n"},
		"doc/guide/two words.md",
	)(
		"doc/guide/index.md",
		strings.NewReader("See [two](two%20words.md#part) there.\n"),
	)
	assertReport(t, "doc/guide/index.md", false, nil, "", l)
}

func TestPointersHeadingEncodedLinkDead(t *testing.T) {
	l := pointer_tester.Headings(
		map[string]string{"doc/guide/two words.md": "# Two\n\n## Part\n"},
		"doc/guide/two words.md",
	)(
		"doc/guide/index.md",
		strings.NewReader("See [two](two%20words.md#parts) there.\n"),
	)
	assertReport(
		t,
		"doc/guide/index.md",
		true,
		pointer_tester.DeadHeading(
			"doc/guide/index.md",
			1,
			"See [two](two%20words.md#parts) there.",
			"Referenced heading does not exist - nearest: #part (\"Part\"), #two (\"Two\")",
		),
		"",
		l,
	)
}
