package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"testing"
)

func TestDetailLinkOSC8(t *testing.T) {
	s := status.New(option.New()).String("a").DetailLink(
		"https://example.org",
		"Jira",
		"",
	)
	assert.String(
		t,
		"a | \x1b]8;;https://example.org\x1b\\Jira\x1b]8;;\x1b\\",
		s.Format(),
	)
}

func TestDetailLinkOSC8DefaultLabel(t *testing.T) {
	s := status.New(option.New()).String("a").DetailLink(
		"https://example.org",
		"",
		"",
	)
	assert.String(
		t,
		"a | \x1b]8;;https://example.org\x1b\\Link\x1b]8;;\x1b\\",
		s.Format(),
	)
}

func TestDetailLinkCopyable(t *testing.T) {
	s := status.New(option.New().Tag(tag.Copyable)).String("a").DetailLink(
		"https://example.org",
		"Jira",
		"",
	)
	assert.String(t, "a\n  https://example.org", s.Format())
}

func TestDetailLinkCopyablePrefix(t *testing.T) {
	s := status.New(option.New().Tag(tag.Copyable)).String("a").DetailLink(
		"https://example.org",
		"Grafana",
		"Prometheus 24h graph",
	)
	assert.String(
		t,
		"a\n  Prometheus 24h graph: https://example.org",
		s.Format(),
	)
}

func TestDetailLinkMarkdown(t *testing.T) {
	s := status.New(option.New().Tag(tag.Markdown)).String("a").DetailLink(
		"https://example.org",
		"Jira",
		"",
	)
	assert.String(
		t,
		"a\n  [Jira](https://example.org)",
		s.Format(),
	)
}

func TestDetailLinkMarkdownPrefix(t *testing.T) {
	s := status.New(option.New().Tag(tag.Markdown)).String("a").DetailLink(
		"https://example.org",
		"Grafana",
		"Prometheus 24h graph",
	)
	assert.String(
		t,
		"a\n  Prometheus 24h graph: [Grafana](https://example.org)",
		s.Format(),
	)
}
