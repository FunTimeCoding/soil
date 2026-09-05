package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "main.go", constant.MainFile)
	assert.String(t, "README.md", constant.ReadmeFile)
}
