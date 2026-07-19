package unit_test

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"strings"
	"testing"
)

func TestChunkGoSourceSplitsAtDeclarations(t *testing.T) {
	body := strings.Repeat("\t// line\n", 200)
	source := fmt.Sprintf(
		"package main\n\nfunc Alpha() {\n%s}\n\nfunc Beta() {\n%s}\n\nfunc Gamma() {\n%s}\n",
		body,
		body,
		body,
	)
	chunks := chunk.Document(source, "main.go")
	assert.Greater(t, 1, float64(len(chunks)))
	assert.StringContains(t, "package main", chunks[0].Text)
	found := false

	for _, c := range chunks {
		if strings.Contains(c.Text, "func Beta") {
			found = true
		}
	}

	if !found {
		t.Errorf("expected a chunk containing func Beta")
	}
}

func TestChunkGoSourceFallsBackOnParseError(t *testing.T) {
	invalid := fmt.Sprintf(
		"package main\n\nfunc broken {{{{\n%s",
		strings.Repeat("text\n", 1000),
	)
	chunks := chunk.Document(invalid, "broken.go")
	assert.Greater(t, 0, float64(len(chunks)))
}

func TestChunkGoSourceShortFile(t *testing.T) {
	source := "package store\n\nfunc Hello() {}\n"
	chunks := chunk.Document(source, "hello.go")
	assert.Count(t, 1, chunks)
	assert.String(t, source, chunks[0].Text)
}
