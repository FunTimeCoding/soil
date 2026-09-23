package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"testing"
)

func TestGroupedLabelsGroupsByObject(t *testing.T) {
	s := newStore(t)
	_, e := s.SetLabel("dcim.device", 7, "owner", "admin")
	errors.PanicOnError(e)
	_, e = s.SetLabel("dcim.device", 7, "backup", "nightly")
	errors.PanicOnError(e)
	_, e = s.SetLabel("dcim.device", 9, "owner", "operations")
	errors.PanicOnError(e)
	_, e = s.SetLabel("virtualization.virtualmachine", 7, "owner", "web")
	errors.PanicOnError(e)
	grouped, f := s.GroupedLabels("dcim.device", []int32{7, 9})
	errors.PanicOnError(f)
	assert.Integer(t, 2, len(grouped))
	assert.Integer(t, 2, len(grouped[7]))
	assert.String(t, "backup", grouped[7][0].Key)
	assert.String(t, "owner", grouped[7][1].Key)
	assert.Integer(t, 1, len(grouped[9]))
	assert.String(t, "operations", grouped[9][0].Value)
}

func TestGroupedLabelsSkipsUnrequestedObjects(t *testing.T) {
	s := newStore(t)
	_, e := s.SetLabel("dcim.device", 7, "owner", "admin")
	errors.PanicOnError(e)
	_, e = s.SetLabel("dcim.device", 9, "owner", "operations")
	errors.PanicOnError(e)
	grouped, f := s.GroupedLabels("dcim.device", []int32{7})
	errors.PanicOnError(f)
	assert.Integer(t, 1, len(grouped))
	assert.Integer(t, 0, len(grouped[9]))
}

func TestGroupedLabelsEmptyRequest(t *testing.T) {
	s := newStore(t)
	_, e := s.SetLabel("dcim.device", 7, "owner", "admin")
	errors.PanicOnError(e)
	grouped, f := s.GroupedLabels("dcim.device", nil)
	errors.PanicOnError(f)
	assert.Integer(t, 0, len(grouped))
}
