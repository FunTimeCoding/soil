package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"testing"
)

func TestVerdictSed(t *testing.T) {
	assert.String(
		t,
		constant.SedMessage,
		guard.Verdict("darwin", "sed -i 's/a/b/' file.go"),
	)
	assert.String(
		t,
		constant.SedMessage,
		guard.Verdict("darwin", "grep x file | sed 's/a/b/'"),
	)
	assert.String(
		t,
		constant.SedMessage,
		guard.Verdict("darwin", "cat file; sed -n 1p file"),
	)
}

func TestVerdictPasses(t *testing.T) {
	assert.String(t, "", guard.Verdict("darwin", "gsed -i 's/a/b/' file.go"))
	assert.String(t, "", guard.Verdict("darwin", "echo parsed"))
	assert.String(t, "", guard.Verdict("darwin", "git grep sedative"))
	assert.String(t, "", guard.Verdict("linux", "sed -i 's/a/b/' file.go"))
}
