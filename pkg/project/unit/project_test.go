package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "main.go", soil.MainFile)
	assert.String(t, "README.md", soil.ReadmeFile)
}
