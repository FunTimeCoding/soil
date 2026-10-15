package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/fixture"
	"testing"
)

func TestChoiceSetSendsValueLabelPairs(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New("netbox.example.org", "token", netbox.WithServer(s.URL))
	defer c.Close()
	v, e := c.CreateChoiceSet("alfa", []string{"bravo", "charlie"})
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		map[string]any{
			"name": "alfa",
			"extra_choices": []any{
				[]any{"bravo", "bravo"},
				[]any{"charlie", "charlie"},
			},
		},
		r.Created,
	)
	assert.Strings(t, []string{"bravo", "charlie"}, v.Choices)
}

func TestSelectionFieldTargetsDevicesAndVirtualMachines(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New("netbox.example.org", "token", netbox.WithServer(s.URL))
	defer c.Close()
	v, e := c.CreateSelectionField("delta", "", "alfa")
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		map[string]any{
			"name": "delta",
			"type": "multiselect",
			"object_types": []any{
				"dcim.device",
				"virtualization.virtualmachine",
			},
			"choice_set": float64(1),
		},
		r.Created,
	)
	assert.String(t, "alfa", v.ChoiceSet)
	assert.String(t, "multiselect", v.Kind)
}
