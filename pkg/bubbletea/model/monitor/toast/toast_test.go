package toast

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestToast(t *testing.T) {
	assert.NotNil(t, New(0, upper.Alfa))
}
