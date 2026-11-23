package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gographd/option"
	"testing"
)

func TestOption(t *testing.T) {
	assert.NotNil(t, option.New())
}
