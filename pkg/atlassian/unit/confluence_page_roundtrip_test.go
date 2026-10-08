package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"testing"
)

func TestRoundTripStoragePlainParagraph(t *testing.T) {
	m := page.ToMarkdown("<p>Hello world</p>")
	assert.String(t, m, page.ToMarkdown(page.ToStorage(m)))
}

func TestRoundTripStorageTwoParagraphs(t *testing.T) {
	m := page.ToMarkdown("<p>First</p><p>Second</p>")
	assert.String(t, m, page.ToMarkdown(page.ToStorage(m)))
}

func TestRoundTripStorageBulletList(t *testing.T) {
	m := page.ToMarkdown("<ul><li><p>Alfa</p></li><li><p>Bravo</p></li></ul>")
	assert.String(t, m, page.ToMarkdown(page.ToStorage(m)))
}

func TestRoundTripStorageHeadingAndParagraph(t *testing.T) {
	m := page.ToMarkdown("<h2>Title</h2><p>Body text</p>")
	assert.String(t, m, page.ToMarkdown(page.ToStorage(m)))
}

func TestRoundTripStorageBoldAndItalic(t *testing.T) {
	m := page.ToMarkdown("<p><strong>bold</strong> and <em>italic</em></p>")
	assert.String(t, m, page.ToMarkdown(page.ToStorage(m)))
}

func TestRoundTripStorageLink(t *testing.T) {
	m := page.ToMarkdown(`<p><a href="https://example.org">Example</a></p>`)
	assert.String(t, m, page.ToMarkdown(page.ToStorage(m)))
}

func TestRoundTripStorageOrderedList(t *testing.T) {
	m := page.ToMarkdown("<ol><li><p>One</p></li><li><p>Two</p></li></ol>")
	assert.String(t, m, page.ToMarkdown(page.ToStorage(m)))
}

func TestRoundTripStorageCodeInline(t *testing.T) {
	m := page.ToMarkdown("<p>Use <code>fmt.Println</code> here</p>")
	assert.String(t, m, page.ToMarkdown(page.ToStorage(m)))
}

func TestRoundTripStorageHardBreak(t *testing.T) {
	m := page.ToMarkdown("<p>Line one<br />Line two</p>")
	assert.String(t, m, page.ToMarkdown(page.ToStorage(m)))
}

func TestRoundTripStorageCodeBlock(t *testing.T) {
	m := page.ToMarkdown(
		`<ac:structured-macro ac:name="code"><ac:plain-text-body><![CDATA[fmt.Println("hello")]]></ac:plain-text-body></ac:structured-macro>`,
	)
	assert.String(t, m, page.ToMarkdown(page.ToStorage(m)))
}

func TestRoundTripStorageNestedList(t *testing.T) {
	m := page.ToMarkdown(
		"<ul><li><p>Outer</p><ul><li><p>Inner</p></li></ul></li></ul>",
	)
	assert.String(t, m, page.ToMarkdown(page.ToStorage(m)))
}

func TestRoundTripStorageTable(t *testing.T) {
	m := page.ToMarkdown(
		"<table><tbody><tr><th><p>Name</p></th><th><p>Value</p></th></tr><tr><td><p>Alfa</p></td><td><p>One</p></td></tr></tbody></table>",
	)
	assert.String(t, m, page.ToMarkdown(page.ToStorage(m)))
}

func TestRoundTripStoragePanelMacro(t *testing.T) {
	m := page.ToMarkdown(
		`<ac:structured-macro ac:name="info"><ac:rich-text-body><p>This is an info panel</p></ac:rich-text-body></ac:structured-macro>`,
	)
	assert.String(t, m, page.ToMarkdown(page.ToStorage(m)))
}

func TestRoundTripMarkdownPlainParagraph(t *testing.T) {
	assert.String(
		t,
		"Hello world",
		page.ToMarkdown(page.ToStorage("Hello world")),
	)
}

func TestRoundTripMarkdownTwoParagraphs(t *testing.T) {
	assert.String(
		t,
		"First\n\nSecond",
		page.ToMarkdown(page.ToStorage("First\n\nSecond")),
	)
}

func TestRoundTripMarkdownBulletList(t *testing.T) {
	assert.String(
		t,
		"- Alfa\n- Bravo",
		page.ToMarkdown(page.ToStorage("- Alfa\n- Bravo")),
	)
}

func TestRoundTripMarkdownHeadingAndParagraph(t *testing.T) {
	assert.String(
		t,
		"## Title\n\nBody text",
		page.ToMarkdown(page.ToStorage("## Title\n\nBody text")),
	)
}

func TestRoundTripMarkdownBoldAndItalic(t *testing.T) {
	assert.String(
		t,
		"**bold** and *italic*",
		page.ToMarkdown(page.ToStorage("**bold** and *italic*")),
	)
}

func TestRoundTripMarkdownLink(t *testing.T) {
	assert.String(
		t,
		"[Example](https://example.org)",
		page.ToMarkdown(page.ToStorage("[Example](https://example.org)")),
	)
}

func TestRoundTripMarkdownOrderedList(t *testing.T) {
	assert.String(
		t,
		"1. One\n2. Two",
		page.ToMarkdown(page.ToStorage("1. One\n2. Two")),
	)
}

func TestRoundTripMarkdownCodeInline(t *testing.T) {
	assert.String(
		t,
		"Use `fmt.Println` here",
		page.ToMarkdown(page.ToStorage("Use `fmt.Println` here")),
	)
}

func TestRoundTripMarkdownHardBreak(t *testing.T) {
	assert.String(
		t,
		"Line one  \nLine two",
		page.ToMarkdown(page.ToStorage("Line one  \nLine two")),
	)
}

func TestRoundTripMarkdownCodeBlock(t *testing.T) {
	assert.String(
		t,
		"```\nfmt.Println(\"hello\")\n```",
		page.ToMarkdown(page.ToStorage("```\nfmt.Println(\"hello\")\n```")),
	)
}

func TestRoundTripMarkdownNestedList(t *testing.T) {
	assert.String(
		t,
		"- Outer\n  \n  - Inner",
		page.ToMarkdown(page.ToStorage("- Outer\n    - Inner")),
	)
}

func TestRoundTripMarkdownTable(t *testing.T) {
	assert.String(
		t,
		"| Name | Value |\n|------|-------|\n| Alfa | One   |",
		page.ToMarkdown(
			page.ToStorage("| Name | Value |\n| --- | --- |\n| Alfa | One |"),
		),
	)
}
