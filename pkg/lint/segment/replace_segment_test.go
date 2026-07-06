package segment

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestReplaceSegmentCamelCase(t *testing.T) {
	assert.String(
		t,
		"directoryName",
		ReplaceSegment("dirName", "dir", "directory"),
	)
}

func TestReplaceSegmentPascalCase(t *testing.T) {
	assert.String(
		t,
		"DirectorySomething",
		ReplaceSegment("DirSomething", "dir", "directory"),
	)
}

func TestReplaceSegmentSingleWord(t *testing.T) {
	assert.String(t, "locator", ReplaceSegment("url", "url", "locator"))
}

func TestReplaceSegmentSingleWordExported(t *testing.T) {
	assert.String(t, "Locator", ReplaceSegment("Url", "url", "locator"))
}

func TestReplaceSegmentMiddle(t *testing.T) {
	assert.String(t, "fooLocator", ReplaceSegment("fooUrl", "url", "locator"))
}

func TestReplaceSegmentMiddleExported(t *testing.T) {
	assert.String(t, "FooLocator", ReplaceSegment("FooUrl", "url", "locator"))
}

func TestReplaceSegmentSnakeCase(t *testing.T) {
	assert.String(
		t,
		"foo_directory",
		ReplaceSegment("foo_dir", "dir", "directory"),
	)
}

func TestReplaceSegmentMultiWordCamel(t *testing.T) {
	assert.String(
		t,
		"modelContextServer",
		ReplaceSegment("mcpServer", "mcp", "model_context"),
	)
}

func TestReplaceSegmentMultiWordPascal(t *testing.T) {
	assert.String(
		t,
		"ModelContextServer",
		ReplaceSegment("McpServer", "mcp", "model_context"),
	)
}

func TestReplaceSegmentMultiWordSnake(t *testing.T) {
	assert.String(
		t,
		"model_context_server",
		ReplaceSegment("mcp_server", "mcp", "model_context"),
	)
}

func TestReplaceSegmentSingleChar(t *testing.T) {
	assert.String(t, "xWrapper", ReplaceSegment("ctxWrapper", "ctx", "x"))
}

func TestReplaceSegmentSingleCharExported(t *testing.T) {
	assert.String(t, "XWrapper", ReplaceSegment("CtxWrapper", "ctx", "x"))
}

func TestReplaceSegmentSingleCharAlone(t *testing.T) {
	assert.String(t, "x", ReplaceSegment("ctx", "ctx", "x"))
}

func TestReplaceSegmentNoMatch(t *testing.T) {
	assert.String(t, "fooBar", ReplaceSegment("fooBar", "baz", "qux"))
}
