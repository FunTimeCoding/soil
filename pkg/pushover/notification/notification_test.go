package notification

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"testing"
)

func TestNotification(t *testing.T) {
	n := New("a", "b", "c")
	n.SetPriority(1)
	assert.String(
		t,
		`{"user":"a","token":"b","message":"c","priority":1}`,
		notation.Encode(n, false),
	)
}
