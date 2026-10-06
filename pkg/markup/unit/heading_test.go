package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/markup/heading"
	"testing"
)

func TestHeadingSlugDropsPunctuation(t *testing.T) {
	assert.String(
		t,
		"pipeline-structure-draft",
		heading.Slug("Pipeline structure (draft)"),
	)
}

func TestHeadingSlugKeepsEverySpace(t *testing.T) {
	assert.String(t, "fish--chips", heading.Slug("Fish & chips"))
}

func TestHeadingSlugKeepsLettersBeyondAscii(t *testing.T) {
	assert.String(t, "größe-und_maß", heading.Slug("Größe und_Maß"))
}

func TestHeadingParseKeepsCodeSpanText(t *testing.T) {
	assert.Any(
		t,
		[]*heading.Heading{heading.New("The goname CLI", "the-goname-cli")},
		heading.Parse("# The `goname` CLI\n"),
	)
}

func TestHeadingParseNumbersDuplicates(t *testing.T) {
	assert.Any(
		t,
		[]*heading.Heading{
			heading.New("Usage", "usage"),
			heading.New("Usage", "usage-1"),
			heading.New("Usage", "usage-2"),
		},
		heading.Parse("## Usage\n\ntext\n\n## Usage\n\n### Usage\n"),
	)
}

func TestHeadingParseSkipsCodeFences(t *testing.T) {
	assert.Count(t, 0, heading.Parse("```\n# not a heading\n```\n"))
}

func TestHeadingParseSkipsFrontMatter(t *testing.T) {
	assert.Any(
		t,
		[]*heading.Heading{heading.New("Title", "title")},
		heading.Parse("---\nbase: pkg/lint\n---\n\n# Title\n"),
	)
}

func TestHeadingParseReadsUnderlinedHeadings(t *testing.T) {
	assert.Any(
		t,
		[]*heading.Heading{
			heading.New("Title", "title"),
			heading.New("Section", "section"),
		},
		heading.Parse("Title\n=====\n\nSection\n-------\n"),
	)
}
