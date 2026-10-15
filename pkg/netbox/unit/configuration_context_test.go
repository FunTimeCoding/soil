package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/netbox/fixture"
	"testing"
)

func TestConfigurationContextReadsPayload(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New("netbox.example.org", "token", netbox.WithServer(s.URL))
	defer c.Close()
	v, e := c.ConfigurationContextByName("alfa")
	assert.FatalOnError(t, e)
	assert.Any(t, map[string]any{"bravo": "charlie"}, v.Payload)
	assert.Integer(t, 1000, v.Weight)
}

func TestClearedAssignmentSendsEveryListEmpty(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New("netbox.example.org", "token", netbox.WithServer(s.URL))
	defer c.Close()
	_, e := c.SetConfigurationContextAssignment(
		"alfa",
		configuration_assignment.New(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)

	for _, k := range []string{
		"roles",
		"clusters",
		"platforms",
		"sites",
		"tenants",
		"tags",
	} {
		assert.Any(t, []any{}, r.Patched[k])
	}
}

func TestPayloadUpdateLeavesAssignmentAlone(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New("netbox.example.org", "token", netbox.WithServer(s.URL))
	defer c.Close()
	_, e := c.SetConfigurationContextPayload(
		"alfa",
		map[string]any{"delta": "echo"},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		map[string]any{"data": map[string]any{"delta": "echo"}},
		r.Patched,
	)
}

func TestUnknownConfigurationContextIsNotFound(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New("netbox.example.org", "token", netbox.WithServer(s.URL))
	defer c.Close()
	_, e := c.ConfigurationContextByName("delta")
	assert.True(t, not_found.Is(e))
}
