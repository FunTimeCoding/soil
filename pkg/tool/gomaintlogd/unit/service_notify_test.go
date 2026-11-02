package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/types/change"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/service_tester"
	"testing"
)

func TestServiceNotifiesOnAdd(t *testing.T) {
	o := service_tester.New(t)
	c := o.Notifier.Subscribe()
	defer o.Notifier.Unsubscribe(c)
	assert.FatalOnError(t, o.Service.Add(sample()))
	assert.True(t, fired(c))
}

func TestServiceNotifiesOnUpdateAndDelete(t *testing.T) {
	o := service_tester.New(t)
	s := o.Service
	v := sample()
	assert.FatalOnError(t, s.Add(v))
	c := o.Notifier.Subscribe()
	defer o.Notifier.Unsubscribe(c)
	user := "bob"
	h := change.New()
	h.User = &user
	result, e := s.Update(v.Identifier, h)
	assert.FatalOnError(t, e)
	assert.String(t, user, result.User)
	assert.True(t, fired(c))
	assert.FatalOnError(t, s.Delete(v.Identifier))
	assert.True(t, fired(c))
}

func TestServiceFailedWriteDoesNotNotify(t *testing.T) {
	o := service_tester.New(t)
	c := o.Notifier.Subscribe()
	defer o.Notifier.Unsubscribe(c)
	o.Close()
	assert.Error(t, o.Service.Add(sample()))
	assert.False(t, fired(c))
}
