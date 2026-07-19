package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"testing"
)

func TestFetch(t *testing.T) {
	assert.NotNil(t, fetch.Command())
}
