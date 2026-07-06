package project

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "main.go", MainFile)
	assert.String(t, "README.md", ReadmeFile)
}
