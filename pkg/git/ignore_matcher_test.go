package git

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestIgnoreMatcher(t *testing.T) {
	root := t.TempDir()

	if e := os.WriteFile(
		filepath.Join(root, ".gitignore"),
		[]byte("/.claude/notes\ntmp/\n"),
		0o600,
	); e != nil {
		t.Fatalf("write: %v", e)
	}

	ignored := IgnoreMatcher(root)
	assert.True(t, ignored(".claude/notes"))
	assert.True(t, ignored(".claude/notes/alfa.md"))
	assert.True(t, ignored("tmp/report.json"))
	assert.False(t, ignored(".claude/skills/bravo/SKILL.md"))
	assert.False(t, ignored("doc/charlie.md"))
}

func TestIgnoreMatcherWithoutRules(t *testing.T) {
	ignored := IgnoreMatcher(t.TempDir())
	assert.False(t, ignored("doc/charlie.md"))
}
