package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/pushover/notification"
	"testing"
)

func TestNotification(t *testing.T) {
	n := notification.New("a", "b", "c")
	n.SetPriority(1)
	assert.String(
		t,
		`{"user":"a","token":"b","message":"c","priority":1}`,
		notation.Encode(n, false),
	)
}
