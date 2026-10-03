package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestACobraToolPrintsTheSameVersionBlockAsEveryTool(t *testing.T) {
	assert.String(
		t,
		"Version: v1.2.3\nGitHash: abc1234\nBuildDate: 2026-01-01T00:00:00Z\nModule: github.com/funtimecoding/soil\nDirty: false\n",
		cobraVersion(t, "--version"),
	)
}

func TestACobraToolPrintsTheVersionAsNotation(t *testing.T) {
	assert.String(
		t,
		"{\n\t\"name\": \"gotest\",\n\t\"version\": \"v1.2.3\",\n\t\"git_hash\": \"abc1234\",\n\t\"build_date\": \"2026-01-01T00:00:00Z\",\n\t\"module\": \"github.com/funtimecoding/soil\",\n\t\"dirty\": false\n}\n",
		cobraVersion(t, "--version", "--notation"),
	)
}
