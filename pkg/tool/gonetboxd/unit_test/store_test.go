package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	s := store.New(lite.NewMemory())
	first, e := s.SetLabel("dcim.device", 7, "owner", "admin")
	errors.PanicOnError(e)
	assert.String(t, "admin", first.Value)
	second, e := s.SetLabel("dcim.device", 7, "owner", "operations")
	errors.PanicOnError(e)
	assert.String(t, "operations", second.Value)
	_, e = s.SetLabel("dcim.device", 7, "backup", "nightly")
	errors.PanicOnError(e)
	_, e = s.SetLabel("virtualization.virtualmachine", 7, "owner", "web")
	errors.PanicOnError(e)
	labels, e := s.Labels("dcim.device", 7)
	errors.PanicOnError(e)
	assert.Integer(t, 2, len(labels))
	assert.String(t, "backup", labels[0].Key)
	assert.String(t, "operations", labels[1].Value)
	errors.PanicOnError(s.RemoveLabel("dcim.device", 7, "backup"))
	assert.NotNil(t, s.RemoveLabel("dcim.device", 7, "backup"))
	labels, e = s.Labels("dcim.device", 7)
	errors.PanicOnError(e)
	assert.Integer(t, 1, len(labels))
}
