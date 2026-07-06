package push

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestPush(t *testing.T) {
	assert.NotNil(t, New(upper.Alfa, 0, false, upper.Bravo))
}
