package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	events "k8s.io/api/events/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestEvent(t *testing.T) {
	assert.NotNil(
		t,
		event.New(
			&events.Event{ObjectMeta: meta.ObjectMeta{Name: upper.Alfa}},
			"",
		),
	)
}
