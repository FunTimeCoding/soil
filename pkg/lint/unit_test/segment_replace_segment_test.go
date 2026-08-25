package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/constant"
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
	assert.String(t, "message", segment.ReplaceSegment("msg", "msg", "message"))
}

func TestReplaceSegmentSingleWordExported(t *testing.T) {
	assert.String(t, "Message", segment.ReplaceSegment("Msg", "msg", "message"))
}

func TestReplaceSegmentMiddle(t *testing.T) {
	assert.String(
		t,
		"fooMessage",
		segment.ReplaceSegment("fooMsg", "msg", "message"),
	)
}

func TestReplaceSegmentMiddleExported(t *testing.T) {
	assert.String(
		t,
		"FooMessage",
		segment.ReplaceSegment("FooMsg", "msg", "message"),
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
	assert.String(
		t,
		"xWrapper",
		segment.ReplaceSegment("ctxWrapper", "ctx", "x"),
	)
}

func TestReplaceSegmentSingleCharExported(t *testing.T) {
	assert.String(
		t,
		"XWrapper",
		segment.ReplaceSegment("CtxWrapper", "ctx", "x"),
	)
}

func TestReplaceSegmentSingleCharAlone(t *testing.T) {
	assert.String(t, "x", segment.ReplaceSegment("ctx", "ctx", "x"))
}

func TestReplaceSegmentNoMatch(t *testing.T) {
	assert.String(t, "fooBar", segment.ReplaceSegment("fooBar", "baz", "qux"))
}

func TestReplaceSegmentTrailingInitialism(t *testing.T) {
	assert.String(
		t,
		"MarshalNotation",
		segment.ReplaceSegment("MarshalJSON", "json", "notation"),
	)
}

func TestReplaceSegmentLeadingInitialism(t *testing.T) {
	assert.String(
		t,
		"LocatorOption",
		segment.ReplaceSegment("URLOption", "url", constant.PointerLocator),
	)
}

func TestReplaceSegmentBareInitialism(t *testing.T) {
	assert.String(
		t,
		"Identifier",
		segment.ReplaceSegment("ID", "id", "identifier"),
	)
}

func TestReplaceSegmentUnexportedInitialism(t *testing.T) {
	assert.String(
		t,
		"nextIdentifier",
		segment.ReplaceSegment("nextID", "id", "identifier"),
	)
}

func TestReplaceSegmentAcronymMultiWord(t *testing.T) {
	assert.String(
		t,
		"StorageDaemonMap",
		segment.ReplaceSegment("OsdMap", "osd", "storage_daemon"),
	)
}

func TestReplaceSegmentUppercaseAcronymMultiWord(t *testing.T) {
	assert.String(
		t,
		"DownStorageDaemon",
		segment.ReplaceSegment("DownOSD", "osd", "storage_daemon"),
	)
}
