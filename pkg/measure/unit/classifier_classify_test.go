package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/measure/classifier"
	"github.com/funtimecoding/soil/pkg/measure/constant"
	"github.com/funtimecoding/soil/pkg/measure/language"
	"testing"
)

func TestClassifyGo(t *testing.T) {
	c := classifier.New(constant.Go).Classify(
		`package main

func main() {
	var n string /*
		comment
		comment
	*/
}
`,
	)
	assert.Integer(t, 1, c.Blank)
	assert.Integer(t, 3, c.Comment)
	assert.Integer(t, 4, c.Code)
}

func TestClassifyGoOneLineBlock(t *testing.T) {
	c := classifier.New(constant.Go).Classify(
		`package main

func main() {
	a := 1 /* one line block */
	/* only comment */
	// line comment
}
`,
	)
	assert.Integer(t, 1, c.Blank)
	assert.Integer(t, 2, c.Comment)
	assert.Integer(t, 4, c.Code)
}

func TestClassifyGoCommentInsideBlock(t *testing.T) {
	c := classifier.New(constant.Go).Classify(
		`package main

/*
// nested line marker
/* second opener is text
*/
func main() {}
`,
	)
	assert.Integer(t, 1, c.Blank)
	assert.Integer(t, 4, c.Comment)
	assert.Integer(t, 2, c.Code)
}

func TestClassifyCodeAfterClose(t *testing.T) {
	c := classifier.New(constant.Go).Classify(
		`/* a
b */ code()
/* c */ /* d */
`,
	)
	assert.Integer(t, 0, c.Blank)
	assert.Integer(t, 2, c.Comment)
	assert.Integer(t, 1, c.Code)
}

func TestClassifyShebangAndHash(t *testing.T) {
	l := language.New("Shell").WithLineComment("#")
	c := classifier.New(l).Classify(
		`#!/bin/sh
# comment

echo hello
`,
	)
	assert.Integer(t, 1, c.Blank)
	assert.Integer(t, 1, c.Comment)
	assert.Integer(t, 2, c.Code)
}

func TestClassifyTripleQuote(t *testing.T) {
	l := language.New("Python").
		WithLineComment("#").
		WithBlockComment(`"""`, `"""`)
	c := classifier.New(l).Classify(
		`#!/bin/python

class A:
	"""comment1
	comment2
	comment3
	"""
	pass
`,
	)
	assert.Integer(t, 1, c.Blank)
	assert.Integer(t, 4, c.Comment)
	assert.Integer(t, 3, c.Code)
}

func TestClassifyNoComments(t *testing.T) {
	c := classifier.New(constant.Markdown).Classify(
		`# Title

Some text // not a comment
`,
	)
	assert.Integer(t, 1, c.Blank)
	assert.Integer(t, 0, c.Comment)
	assert.Integer(t, 2, c.Code)
}

func TestClassifyByteOrderMarkAndNoTrailingNewline(t *testing.T) {
	c := classifier.New(constant.Go).Classify("\xef\xbb\xbf// only\ncode")
	assert.Integer(t, 0, c.Blank)
	assert.Integer(t, 1, c.Comment)
	assert.Integer(t, 1, c.Code)
	assert.Integer(t, 2, c.Total())
}

func TestClassifyEmpty(t *testing.T) {
	c := classifier.New(constant.Go).Classify("")
	assert.Integer(t, 0, c.Total())
}
