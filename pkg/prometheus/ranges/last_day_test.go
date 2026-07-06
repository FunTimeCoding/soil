package ranges

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
	"time"
)

func TestLastDay(t *testing.T) {
	assert.Any(t, time.Minute, LastDay(1).Step)
}
