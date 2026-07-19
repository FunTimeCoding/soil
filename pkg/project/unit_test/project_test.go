package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/project"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "main.go", project.MainFile)
	assert.String(t, "README.md", project.ReadmeFile)
}
