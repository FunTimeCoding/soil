package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/unit/kind_tester"
	"testing"
)

func TestMarkedAfterMarkForSameNameAndBuild(t *testing.T) {
	directory := kind_tester.Module(t)
	store := t.TempDir()
	first := index.New(store, directory).WithFingerprint("alfa")
	assert.False(t, first.Marked("bravo", "charlie"))
	first.Mark("bravo", "charlie")
	again := index.New(store, directory).WithFingerprint("alfa")
	assert.True(t, again.Marked("bravo", "charlie"))
	assert.False(t, again.Marked("delta", "charlie"))
	assert.False(t, again.Marked("bravo", "echo"))
	other := index.New(store, directory).WithFingerprint("foxtrot")
	assert.False(t, other.Marked("bravo", "charlie"))
}
