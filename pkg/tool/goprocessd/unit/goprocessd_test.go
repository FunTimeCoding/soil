package unit

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/environment"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/procfile"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/socket"
	"sort"
	"strings"
	"testing"
)

func TestBuildMergesBaseAndOverlay(t *testing.T) {
	e := environment.New([]string{"A=1", "B=2"})
	e.Set("C", "3")
	result := e.Build()
	sort.Strings(result)
	assert.Integer(t, 3, len(result))
	assert.String(t, "A=1", result[0])
	assert.String(t, "B=2", result[1])
	assert.String(t, "C=3", result[2])
}

func TestBuildOverlayOverridesBase(t *testing.T) {
	e := environment.New([]string{"A=1", "B=2"})
	e.Set("B", "changed")
	result := e.Build()
	sort.Strings(result)
	assert.Integer(t, 2, len(result))
	assert.String(t, "A=1", result[0])
	assert.String(t, "B=changed", result[1])
}

func TestBuildEmptyOverlay(t *testing.T) {
	e := environment.New([]string{"A=1"})
	result := e.Build()
	assert.Integer(t, 1, len(result))
	assert.String(t, "A=1", result[0])
}

func TestSplitNull(t *testing.T) {
	input := "KEY=value\x00OTHER=two\x00"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(environment.SplitNull)
	var tokens []string

	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	assert.Integer(t, 2, len(tokens))
	assert.String(t, "KEY=value", tokens[0])
	assert.String(t, "OTHER=two", tokens[1])
}

func TestSplitNullEmbeddedNewline(t *testing.T) {
	input := "KEY=line1\nline2\x00OTHER=value\x00"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(environment.SplitNull)
	var tokens []string

	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	assert.Integer(t, 2, len(tokens))
	assert.String(t, "KEY=line1\nline2", tokens[0])
	assert.String(t, "OTHER=value", tokens[1])
}

func TestSplitNullNoTrailingNull(t *testing.T) {
	input := "KEY=value"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(environment.SplitNull)
	var tokens []string

	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	assert.Integer(t, 1, len(tokens))
	assert.String(t, "KEY=value", tokens[0])
}

func TestParseValidEntries(t *testing.T) {
	path := writeProcfile(
		t,
		"goansibled: go run cmd/goansibled/main.go\ngoclauded: go run cmd/goclauded/main.go\n",
	)
	entries, e := procfile.Parse(path)
	errors.PanicOnError(e)
	assert.Integer(t, 2, len(entries))
	assert.String(t, "goansibled", entries[0].Name)
	assert.String(t, "go run cmd/goansibled/main.go", entries[0].Command)
	assert.String(t, "goclauded", entries[1].Name)
	assert.String(t, "go run cmd/goclauded/main.go", entries[1].Command)
}

func TestParseSkipsCommentsAndBlankLines(t *testing.T) {
	path := writeProcfile(
		t,
		"# comment\n\ngoansibled: go run cmd/goansibled/main.go\n\n# another comment\n",
	)
	entries, e := procfile.Parse(path)
	errors.PanicOnError(e)
	assert.Integer(t, 1, len(entries))
	assert.String(t, "goansibled", entries[0].Name)
}

func TestParseTrimsWhitespace(t *testing.T) {
	path := writeProcfile(
		t,
		"  goansibled  :  go run cmd/goansibled/main.go  \n",
	)
	entries, e := procfile.Parse(path)
	errors.PanicOnError(e)
	assert.String(t, "goansibled", entries[0].Name)
	assert.String(t, "go run cmd/goansibled/main.go", entries[0].Command)
}

func TestParseEmptyFileReturnsError(t *testing.T) {
	path := writeProcfile(t, "")
	_, e := procfile.Parse(path)
	assert.True(t, e != nil)
}

func TestParseCommandWithColons(t *testing.T) {
	path := writeProcfile(t, "web: sh -c \"echo host:port\"\n")
	entries, e := procfile.Parse(path)
	errors.PanicOnError(e)
	assert.String(t, "sh -c \"echo host:port\"", entries[0].Command)
}

func TestPathDeterministic(t *testing.T) {
	a := socket.Path("/tmp/project/Procfile")
	b := socket.Path("/tmp/project/Procfile")
	assert.String(t, a, b)
}

func TestPathDifferentDirectories(t *testing.T) {
	a := socket.Path("/tmp/alfa/Procfile")
	b := socket.Path("/tmp/bravo/Procfile")
	assert.True(t, a != b)
}

func TestPathEndsWith(t *testing.T) {
	result := socket.Path("/tmp/project/Procfile")
	assert.True(t, strings.HasSuffix(result, ".sock"))
}
