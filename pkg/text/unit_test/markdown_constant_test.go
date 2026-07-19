package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text/markdown/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "WIKI_PATH", constant.WikiPathEnvironment)
}
