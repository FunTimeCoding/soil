package transcript_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
	"testing"
)

func Write(
	t *testing.T,
	lines ...string,
) string {
	t.Helper()
	result := filepath.Join(t.TempDir(), "session.jsonl")
	assert.FatalOnError(
		t,
		os.WriteFile(result, []byte(join.NewLine(lines)), 0o600),
	)

	return result
}
