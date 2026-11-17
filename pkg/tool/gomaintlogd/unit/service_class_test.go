package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/types/change"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/service_tester"
	"testing"
)

func TestServiceDeleteOfUnknownEntryIsNotFound(t *testing.T) {
	assert.True(t, not_found.Is(service_tester.New(t).Service.Delete(9999)))
}

func TestServiceRefusesEntryWithoutAction(t *testing.T) {
	o := service_tester.New(t)
	c := o.Notifier.Subscribe()
	defer o.Notifier.Unsubscribe(c)
	v := service_tester.Sample()
	v.Action = ""
	assert.True(t, validation.Is(o.Service.Add(v)))
	assert.NotSignalled(t, c)
}

func TestServiceRefusesUpdateClearingUser(t *testing.T) {
	o := service_tester.New(t)
	s := o.Service
	v := service_tester.Sample()
	assert.FatalOnError(t, s.Add(v))
	c := o.Notifier.Subscribe()
	defer o.Notifier.Unsubscribe(c)
	empty := ""
	h := change.New()
	h.User = &empty
	_, e := s.Update(v.Identifier, h)
	assert.True(t, validation.Is(e))
	assert.NotSignalled(t, c)
	stored, f := s.Entry(v.Identifier)
	assert.FatalOnError(t, f)
	assert.String(t, "alice", stored.User)
}
