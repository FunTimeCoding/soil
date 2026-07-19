package chunk

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"strings"
	"testing"
)

func TestDocumentShortContent(t *testing.T) {
	content := "# Short\n\nThis is a short document.\n"
	chunks := Document(content, "test.md")
	assert.Count(t, 1, chunks)
	assert.String(t, content, chunks[0].Text)
	assert.Integer(t, 0, chunks[0].Position)
}

func TestDocumentSplitsAtHeadings(t *testing.T) {
	sections := make([]string, 5)

	for i := range sections {
		sections[i] = strings.Repeat("x", constant.ChunkSize/4)
	}

	content := fmt.Sprintf(
		"# First\n\n%s\n# Second\n\n%s\n# Third\n\n%s\n# Fourth\n\n%s\n# Fifth\n\n%s",
		sections[0],
		sections[1],
		sections[2],
		sections[3],
		sections[4],
	)
	chunks := Document(content, "test.md")
	assert.Greater(t, 1, float64(len(chunks)))
	assert.StringContains(t, "# First", chunks[0].Text)
}

func TestDocumentPreservesCodeFences(t *testing.T) {
	before := strings.Repeat("text ", constant.ChunkSize/8)
	code := strings.Repeat("line\n", 20)
	after := strings.Repeat("more ", constant.ChunkSize/8)
	content := fmt.Sprintf(
		"# Before\n\n%s\n```go\n%s```\n\n# After\n\n%s",
		before,
		code,
		after,
	)
	chunks := Document(content, "test.md")

	for _, c := range chunks {
		opens := strings.Count(c.Text, "```")

		if opens == 1 {
			t.Errorf("chunk splits inside code fence")
		}
	}
}
