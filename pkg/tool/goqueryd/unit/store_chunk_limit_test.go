package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestDocumentCutsLongLineOnRuneBoundary(t *testing.T) {
	chunks := chunk.Document(
		join.Empty("# Wide\n\n", strings.Repeat("あいうえお", 400)),
		"wide.md",
		service_tester.New(t).Reranker,
	)
	assert.Greater(t, 1, len(chunks))

	for _, c := range chunks {
		assert.True(t, utf8.ValidString(c.Text))
	}
}

func TestDocumentStartsOverlapOnRuneBoundary(t *testing.T) {
	chunks := chunk.Document(
		strings.Repeat(join.Empty(strings.Repeat("あ", 33), "\n"), 80),
		"lines.md",
		service_tester.New(t).Reranker,
	)
	assert.Greater(t, 1, len(chunks))

	for _, c := range chunks {
		assert.True(t, utf8.ValidString(c.Text))
	}
}
