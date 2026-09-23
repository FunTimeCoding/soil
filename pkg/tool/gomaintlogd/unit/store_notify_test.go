package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"testing"
)

func TestStoreNotifiesOnAdd(t *testing.T) {
	n := notifier.New()
	s := newStore(t, n)
	c := n.Subscribe()
	defer n.Unsubscribe(c)
	assert.FatalOnError(t, s.Add(sample()))
	assert.True(t, fired(c))
}

func TestStoreNotifiesOnUpdateAndDelete(t *testing.T) {
	n := notifier.New()
	s := newStore(t, n)
	v := sample()
	assert.FatalOnError(t, s.Add(v))
	c := n.Subscribe()
	defer n.Unsubscribe(c)
	v.User = "bob"
	assert.FatalOnError(t, s.Update(v))
	assert.True(t, fired(c))
	assert.FatalOnError(t, s.Delete(v.Identifier))
	assert.True(t, fired(c))
}
