package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gohook/job"
	"testing"
)

func TestCommandSubstitutesArguments(t *testing.T) {
	j := job.New("gocommit check {1}")
	assert.String(
		t,
		"gocommit check .git/COMMIT_EDITMSG",
		j.Command([]string{".git/COMMIT_EDITMSG"}),
	)
}

func TestCommandSubstitutesAll(t *testing.T) {
	j := job.New("echo {@}")
	assert.String(t, "echo a b", j.Command([]string{"a", "b"}))
	assert.String(t, "echo ", j.Command(nil))
}

func TestCommandWithoutPlaceholders(t *testing.T) {
	assert.String(t, "task lint", job.New("task lint").Command([]string{"x"}))
}

func TestCommandMissingArgumentStaysLiteral(t *testing.T) {
	assert.String(t, "echo {2}", job.New("echo {2}").Command([]string{"a"}))
}
