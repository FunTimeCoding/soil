package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"strings"
	"testing"
)

func TestChunkGoSourceSplitsAtDeclarations(t *testing.T) {
	body := strings.Repeat("\t// line\n", 200)
	source := fmt.Sprintf(
		"package main\n\nfunc Alfa() {\n%s}\n\nfunc Bravo() {\n%s}\n\nfunc Charlie() {\n%s}\n",
		body,
		body,
		body,
	)
	chunks := chunk.Document(source, constant.MainFile, mock_reranker.New())
	assert.Greater(t, 1, len(chunks))
	assert.StringContains(t, "package main", chunks[0].Text)
	found := false

	for _, c := range chunks {
		if strings.Contains(c.Text, "func Bravo") {
			found = true
		}
	}

	if !found {
		t.Errorf("expected a chunk containing func Bravo")
	}
}

func TestChunkGoSourceFallsBackOnParseError(t *testing.T) {
	invalid := fmt.Sprintf(
		"package main\n\nfunc broken {{{{\n%s",
		strings.Repeat("text\n", 1000),
	)
	chunks := chunk.Document(invalid, "broken.go", mock_reranker.New())
	assert.Greater(t, 0, len(chunks))
}

func TestChunkGoSourceShortFile(t *testing.T) {
	source := "package store\n\nfunc Hello() {}\n"
	chunks := chunk.Document(source, "hello.go", mock_reranker.New())
	assert.Count(t, 1, chunks)
	assert.String(t, source, chunks[0].Text)
}

func TestChunkGoSourceSplitsAShortFileOverTheAllowance(t *testing.T) {
	r := mock_reranker.New()
	defer errors.PanicClose(r)
	r.SetAllowance(10)
	chunks := chunk.Document(
		"package store\n\nfunc Alfa() {\n\treturn\n}\n\nfunc Bravo() {\n\treturn\n}\n\nfunc Charlie() {\n\treturn\n}\n",
		"short.go",
		r,
	)
	assert.Count(t, 4, chunks)
	assert.String(t, "package store\n\n", chunks[0].Text)
	assert.String(t, "func Alfa() {\n\treturn\n}\n\n", chunks[1].Text)
	assert.String(t, "func Bravo() {\n\treturn\n}\n\n", chunks[2].Text)
	assert.String(t, "func Charlie() {\n\treturn\n}\n", chunks[3].Text)
}

func TestChunkGoSourceKeepsAShortFileAtTheAllowance(t *testing.T) {
	source := "package store\n\nfunc Alfa() {\n\treturn\n}\n\nfunc Bravo() {\n\treturn\n}\n\nfunc Charlie() {\n\treturn\n}\n"
	r := mock_reranker.New()
	defer errors.PanicClose(r)
	r.SetAllowance(17)
	chunks := chunk.Document(source, "short.go", r)
	assert.Count(t, 1, chunks)
	assert.String(t, source, chunks[0].Text)
}
