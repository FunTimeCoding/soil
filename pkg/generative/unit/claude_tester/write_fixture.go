package claude_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"path/filepath"
	"testing"
)

func WriteFixture(
	t *testing.T,
	lines []string,
) *claude.Client {
	t.Helper()
	base := t.TempDir()
	f := system.Create(filepath.Join(base, "fixture.jsonl"))

	for _, line := range lines {
		writer.Print(f, "%s\n", line)
	}

	errors.PanicClose(f)

	return claude.NewDirectory(base)
}
