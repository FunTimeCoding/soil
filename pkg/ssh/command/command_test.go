package command

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestCommand(t *testing.T) {
	assert.NotNil(t, New(upper.Alfa))
}
