package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "---", constant.FrontMatterDelimiter)
}
