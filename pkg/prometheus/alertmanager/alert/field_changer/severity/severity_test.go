package severity

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestSeverity(t *testing.T) {
	assert.NotNil(t, New(upper.Alfa, upper.Bravo, upper.Charlie))
}
