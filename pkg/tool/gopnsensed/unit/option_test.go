package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/option"
	"testing"
)

func TestOptionNew(t *testing.T) {
	assert.NotNil(t, option.New())
}
