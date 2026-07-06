package ranges

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
	"time"
)

func TestLastHour(t *testing.T) {
	assert.Any(t, time.Minute, LastHour(1).Step)
}
