package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/procfile"
	"os"
	"path/filepath"
	"testing"
)

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

func writeProcfile(
	t *testing.T,
	content string,
) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "Procfile")
	errors.PanicOnError(os.WriteFile(path, []byte(content), 0644))

	return path
}
