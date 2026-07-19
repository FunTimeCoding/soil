package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/toast"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestToast(t *testing.T) {
	assert.NotNil(t, toast.New(0, upper.Alfa))
}
