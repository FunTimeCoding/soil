package page

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestRoundTripStorageToMarkdownToStorage(t *testing.T) {
	cases := []struct {
		name    string
		storage string
	}{
		{
			"plain paragraph",
			"<p>Hello world</p>",
		},
		{
			"two paragraphs",
			"<p>First</p><p>Second</p>",
		},
		{
			"bullet list",
			"<ul><li><p>Alfa</p></li><li><p>Bravo</p></li></ul>",
		},
		{
			"heading and paragraph",
			"<h2>Title</h2><p>Body text</p>",
		},
		{
			"bold and italic",
			"<p><strong>bold</strong> and <em>italic</em></p>",
		},
		{
			"link",
			`<p><a href="https://example.org">Example</a></p>`,
		},
		{
			"ordered list",
			"<ol><li><p>One</p></li><li><p>Two</p></li></ol>",
		},
		{
			"code inline",
			"<p>Use <code>fmt.Println</code> here</p>",
		},
		{
			"hard break",
			"<p>Line one<br />Line two</p>",
		},
		{
			"code block",
			`<ac:structured-macro ac:name="code"><ac:plain-text-body><![CDATA[fmt.Println("hello")]]></ac:plain-text-body></ac:structured-macro>`,
		},
		{
			"nested list",
			"<ul><li><p>Outer</p><ul><li><p>Inner</p></li></ul></li></ul>",
		},
		{
			"table",
			"<table><tbody><tr><th><p>Name</p></th><th><p>Value</p></th></tr><tr><td><p>Alfa</p></td><td><p>One</p></td></tr></tbody></table>",
		},
		{
			"confluence info macro",
			`<ac:structured-macro ac:name="info"><ac:rich-text-body><p>This is an info panel</p></ac:rich-text-body></ac:structured-macro>`,
		},
	}

	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				markdown := ToMarkdown(c.storage)
				restored := ToStorage(markdown)

				// Round-trip the restored storage back to markdown
				// to verify the semantic content survives
				assert.String(t, markdown, ToMarkdown(restored))
			},
		)
	}
}

func TestRoundTripMarkdownToStorageToMarkdown(t *testing.T) {
	cases := []struct {
		name     string
		markdown string
		expect   string // empty means expect identical round-trip
	}{
		{
			"plain paragraph",
			"Hello world",
			"",
		},
		{
			"two paragraphs",
			"First\n\nSecond",
			"",
		},
		{
			"bullet list",
			"- Alfa\n- Bravo",
			"",
		},
		{
			"heading and paragraph",
			"## Title\n\nBody text",
			"",
		},
		{
			"bold and italic",
			"**bold** and *italic*",
			"",
		},
		{
			"link",
			"[Example](https://example.org)",
			"",
		},
		{
			"ordered list",
			"1. One\n2. Two",
			"",
		},
		{
			"code inline",
			"Use `fmt.Println` here",
			"",
		},
		{
			"hard break",
			"Line one  \nLine two",
			"",
		},
		{
			"code block",
			"```\nfmt.Println(\"hello\")\n```",
			"",
		},
		{
			// drift: blank line inserted between outer and inner item
			"nested list",
			"- Outer\n    - Inner",
			"- Outer\n  \n  - Inner",
		},
		{
			// drift: separator and cell padding normalized by html-to-markdown
			"table",
			"| Name | Value |\n| --- | --- |\n| Alfa | One |",
			"| Name | Value |\n|------|-------|\n| Alfa | One   |",
		},
	}

	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				restored := ToMarkdown(ToStorage(c.markdown))
				expect := c.markdown

				if c.expect != "" {
					expect = c.expect
				}

				assert.String(t, expect, restored)
			},
		)
	}
}
