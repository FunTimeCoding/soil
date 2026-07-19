package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"testing"
)

func TestReplaceSegmentCamelCase(t *testing.T) {
	assert.String(
		t,
		"directoryName",
		segment.ReplaceSegment("dirName", "dir", "directory"),
	)
}

func TestReplaceSegmentPascalCase(t *testing.T) {
	assert.String(
		t,
		"DirectorySomething",
		segment.ReplaceSegment("DirSomething", "dir", "directory"),
	)
}

func TestReplaceSegmentSingleWord(t *testing.T) {
	assert.String(t, "locator", segment.ReplaceSegment("url", "url", "locator"))
}

func TestReplaceSegmentSingleWordExported(t *testing.T) {
	assert.String(t, "Locator", segment.ReplaceSegment("Url", "url", "locator"))
}

func TestReplaceSegmentMiddle(t *testing.T) {
	assert.String(
		t,
		"fooLocator",
		segment.ReplaceSegment(
			"fooUrl",
			"url",
			"locator",
		),
	)
}

func TestReplaceSegmentMiddleExported(t *testing.T) {
	assert.String(
		t,
		"FooLocator",
		segment.ReplaceSegment(
			"FooUrl",
			"url",
			"locator",
		),
	)
}

func TestReplaceSegmentSnakeCase(t *testing.T) {
	assert.String(
		t,
		"foo_directory",
		segment.ReplaceSegment("foo_dir", "dir", "directory"),
	)
}

func TestReplaceSegmentMultiWordCamel(t *testing.T) {
	assert.String(
		t,
		"modelContextServer",
		segment.ReplaceSegment("mcpServer", "mcp", "model_context"),
	)
}

func TestReplaceSegmentMultiWordPascal(t *testing.T) {
	assert.String(
		t,
		"ModelContextServer",
		segment.ReplaceSegment("McpServer", "mcp", "model_context"),
	)
}

func TestReplaceSegmentMultiWordSnake(t *testing.T) {
	assert.String(
		t,
		"model_context_server",
		segment.ReplaceSegment("mcp_server", "mcp", "model_context"),
	)
}

func TestReplaceSegmentSingleChar(t *testing.T) {
	assert.String(t, "xWrapper", segment.ReplaceSegment("ctxWrapper", "ctx", "x"))
}

func TestReplaceSegmentSingleCharExported(t *testing.T) {
	assert.String(t, "XWrapper", segment.ReplaceSegment("CtxWrapper", "ctx", "x"))
}

func TestReplaceSegmentSingleCharAlone(t *testing.T) {
	assert.String(t, "x", segment.ReplaceSegment("ctx", "ctx", "x"))
}

func TestReplaceSegmentNoMatch(t *testing.T) {
	assert.String(t, "fooBar", segment.ReplaceSegment("fooBar", "baz", "qux"))
}
