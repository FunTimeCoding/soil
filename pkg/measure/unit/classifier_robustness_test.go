package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/measure/classifier"
	"github.com/funtimecoding/soil/pkg/measure/constant"
	"testing"
)

func TestClassifyOpenerInsideString(t *testing.T) {
	c := classifier.New(constant.Go).Classify(
		`package main

func main() {
	st := "/*"
	a := 1
	en := "*/"
	/* comment */
}
`,
	)
	assert.Integer(t, 1, c.Blank)
	assert.Integer(t, 1, c.Comment)
	assert.Integer(t, 6, c.Code)
}

func TestClassifyEscapedQuote(t *testing.T) {
	c := classifier.New(constant.Go).Classify(
		`s := "a \" /* b"
// comment
`,
	)
	assert.Integer(t, 1, c.Comment)
	assert.Integer(t, 1, c.Code)
}

func TestClassifyRawStringSpansLines(t *testing.T) {
	c := classifier.New(constant.Go).Classify(
		"s := `\n// not a comment\n/* nor this\n`\n// real\n",
	)
	assert.Integer(t, 1, c.Comment)
	assert.Integer(t, 4, c.Code)
}

func TestClassifyUnterminatedQuoteResetsPerLine(t *testing.T) {
	c := classifier.New(constant.Go).Classify(
		`s := "broken
// comment
`,
	)
	assert.Integer(t, 1, c.Comment)
	assert.Integer(t, 1, c.Code)
}

func TestClassifyLineCommentInsideString(t *testing.T) {
	c := classifier.New(constant.Go).Classify(
		`s := "http://example" // trailing
`,
	)
	assert.Integer(t, 0, c.Comment)
	assert.Integer(t, 1, c.Code)
}

func TestClassifyNestedBlock(t *testing.T) {
	c := classifier.New(constant.Rust).Classify(
		`/* outer
/* inner */
still comment */
fn main() {}
`,
	)
	assert.Integer(t, 3, c.Comment)
	assert.Integer(t, 1, c.Code)
}

func TestClassifyNotNestedByDefault(t *testing.T) {
	c := classifier.New(constant.C).Classify(
		`/* outer
/* inner */
now code */
`,
	)
	assert.Integer(t, 2, c.Comment)
	assert.Integer(t, 1, c.Code)
}

func TestClassifySingleQuoteInsideComment(t *testing.T) {
	c := classifier.New(constant.Shell).Classify(
		`# don't
echo 'x'
`,
	)
	assert.Integer(t, 1, c.Comment)
	assert.Integer(t, 1, c.Code)
}
