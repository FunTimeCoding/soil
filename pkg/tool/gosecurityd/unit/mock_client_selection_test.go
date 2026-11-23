package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"testing"
)

func TestMockSelectsSystemsLikeDaemon(t *testing.T) {
	m := mock_client.New(t)
	x := t.Context()

	for name, labels := range map[string]map[string]string{
		"quebec": {"group": "red"},
		"romeo":  {"group": "blue"},
	} {
		assert.FatalOnError(
			t,
			m.PutSystem(
				x,
				name,
				client.SystemSnapshot{Release: "kilo", Labels: &labels},
			),
		)
	}

	red := []string{"group:red"}
	labeled, e := m.ListSystems(x, &client.ListSystemsParams{Label: &red})
	assert.FatalOnError(t, e)
	assert.Count(t, 1, labeled)
	assert.String(t, "quebec", labeled[0].Name)
	part := "ROM"
	found, f := m.ListSystems(x, &client.ListSystemsParams{Search: &part})
	assert.FatalOnError(t, f)
	assert.Count(t, 1, found)
	assert.String(t, "romeo", found[0].Name)
	supported := constant.ShowSupported
	none, g := m.ListSystems(x, &client.ListSystemsParams{Show: &supported})
	assert.FatalOnError(t, g)
	assert.Count(t, 0, none)
	malformed := []string{"group"}
	_, h := m.ListSystems(x, &client.ListSystemsParams{Label: &malformed})
	assert.True(t, validation.Is(h))
}
