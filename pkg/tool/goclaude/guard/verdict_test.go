package guard

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"testing"
)

func TestVerdictSed(t *testing.T) {
	assert.String(
		t,
		constant.SedMessage,
		verdict("darwin", "sed -i 's/a/b/' file.go"),
	)
	assert.String(
		t,
		constant.SedMessage,
		verdict("darwin", "grep x file | sed 's/a/b/'"),
	)
	assert.String(
		t,
		constant.SedMessage,
		verdict("darwin", "cat file; sed -n 1p file"),
	)
}

func TestVerdictPasses(t *testing.T) {
	assert.String(t, "", verdict("darwin", "gsed -i 's/a/b/' file.go"))
	assert.String(t, "", verdict("darwin", "echo parsed"))
	assert.String(t, "", verdict("darwin", "git grep sedative"))
	assert.String(t, "", verdict("linux", "sed -i 's/a/b/' file.go"))
}
